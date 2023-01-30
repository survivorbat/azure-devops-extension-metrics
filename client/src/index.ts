import { v4 as uuidv4 } from 'uuid';
import { message } from './message/input';
import * as net from 'net';

export interface Config {
  host: string;
  port: number;
  errorCallback: (err: Error) => void;
}

const sendMessage = async (id: string, config: Config): Promise<void> => {
  const input = message.Input.fromObject({ id, timestamp: Date.now() });
  const data = input.serialize();

  const client = new net.Socket();

  client.connect(config.port, config.host, () => {
    client.write(data);
    client.destroy();
  });

  client.on('error', (err) => {
    config.errorCallback(err);
  });
};

export const start = async (config: Config): Promise<() => Promise<void>> => {
  // ID for both messages
  const id = uuidv4();

  // Initial message
  await sendMessage(id, config);

  // Finalization message
  return (): Promise<void> => sendMessage(id, config);
};

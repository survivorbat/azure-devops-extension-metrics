import { v4 as uuidv4 } from 'uuid';
import { message } from './message/input';
import * as net from 'net';

export interface Config {
  host: string;
  port: number;
}

const sendMessage = async (id: string, config: Config): Promise<void> => {
  const input = message.Input.fromObject({ id, timestamp: Date.now() });
  const data = input.serialize();

  const client = new net.Socket();

  client.connect(config.port, config.host, () => {
    client.write(data);
  });
};


export const start = async (config: Config): Promise<() => Promise<void>> => {
  const id = uuidv4();

  await sendMessage(id, config);

  return async () => sendMessage(id, config);
};

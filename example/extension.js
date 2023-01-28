const start = require('../client/dist').start;

const config = {
    host: 'localhost',
    port: 8080,
}

const run = async () => {
    const end = await start(config)

    console.log('Extension is running')

    await end();
}

run().then()
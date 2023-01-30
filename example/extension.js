const start = require('../client/dist').start;

const config  = {
    host: 'localhost',
    port: 3000,
    errorCallback: console.error,
}

const run = async () => {
    const end = await start(config)

    console.log('Extension is running')

    setTimeout(() => {
        console.log('Extension is stopping')
        return end();
    }, 2000)
}

run().then()
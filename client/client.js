const _ = require('underscore');
const async = require('async');
const net = require('net');
const readline = require('readline');

module.exports = class Client {
  constructor(port) {
    this._port = port;
  }

  run(commands, done) {
    function reject(msg) {
      s.destroy();
      done(msg);
    }

    done = _.once(done);

    var s = net.connect({ port: this._port }),
      rl = readline.createInterface({ input: s }),
      handshaked = false,
      output = [];

    s.on('error', done);
    s.on('connect', () => {
      console.log("-->client connected to server")

    });

    rl.on('line', line => {
      console.log('line in')
      if (handshaked) {
        output.push(line);

        return;
      }

      console.log('handhake!!!!')
      handshaked = true;

      if (line !== 'hello')
        return reject('Expected: hello');

      s.write(commands.map(c => c + '\r\n').join('  '));
      s.write('quit\r\n');
    });

    rl.on('close', () => done(null, output));

    setTimeout(() => done('timeout'), 1000);
  }
}

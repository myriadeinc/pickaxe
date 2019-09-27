'use strict';

const path = require('path');
const root = path.resolve(path.join(__dirname, '..', '..'));

const config = require('nconf')
    .argv()
    .env({lowerCase: true, separator: '__'})
    .file('environment', {
      file: path.join(root, 'config', `${process.env.NODE_ENV}.json`),
    })
    .file('defaults', {
      file: path.join(root, 'config', 'default.json'),
    });

const Config = {

  get: config.get.bind(config),

  environment: process.env.NODE_ENV,

  isProduction: () => {
    return Config.environment === 'production';
  },
};

module.exports = Config;

'use strict';
const Err = require('egads')
    .extend('Unexpected error occured', 500, 'InternalServerError');

Err.Account = Err.extend('Account Error', 500, 'AccountError');

module.exports = Err;

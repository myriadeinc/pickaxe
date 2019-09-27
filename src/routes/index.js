'use strict';
/*eslint-disable */
const router = require('express').Router();
/* eslint-enable */

router.use('/account', require('src/routes/account.router.js'));
router.use('/email', require('src/routes/email.router.js'));

module.exports = router;

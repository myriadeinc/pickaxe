const path = require('path');

const rootPath = path.resolve(`${__dirname}/..`);
require('app-module-path').addPath(rootPath);

const config = require('src/util/config.js');
const logger = require('src/util/logger.js');
const db = require('src/util/db.js');

let server;

const start = async () => {
  logger.core.info('Starting Diamond: Identity Service for Myriade');
  logger.core.info('Initializing database.');

  await db.init(
      config.get('db'),
      logger.db
  );
  logger.core.info('Database initialized.');

  const port = config.get('port');
  const app = require('./app');
  server = app.listen(port, () => {
    logger.core.info(`Service started on port ${port}`);
  });
};

const gracefulShutdown = () => {
  server.close(async () => {
    logger.core.info('Gracefully closing the app');
    cache.close();
    try {
      await db.close();
    } catch (err) {
      logger.core.warn(`Failed to close the DB Connection - Error: ${err}`);
    } finally {
      process.exit(0);
    }
  });
};

process.on('SIGTERM', gracefulShutdown);
process.on('SIGINT', gracefulShutdown);

process.on('unhandledRejection', (err) => {
  logger.core.error('Unhandled promise rejection!', err);
  process.exit(1);
});

start()
    .catch((err) => {
      logger.core.error(err);
      process.exit(1);
    });

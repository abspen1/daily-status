"use strict";
const fs = require("fs");
const Redis = require("ioredis");

module.exports = async (event, context) => {
  const pass = fs.readFileSync("/var/openfaas/secrets/redis-password", "utf8");

  const redis = new Redis({ post: 6379, host: "192.168.1.6", password: pass });
  var reply = await redis.hgetall("daily_status");

  context.succeed(reply);
};

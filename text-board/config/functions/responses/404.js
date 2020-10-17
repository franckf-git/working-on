'use strict';
var http = require('http'),
  fs = require('fs');
const { pathToFileURL } = require('url');

module.exports = async (ctx) => {
  return ctx.redirect('/')
};

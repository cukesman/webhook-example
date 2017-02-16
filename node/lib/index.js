'use strict';

const Joi = require('joi');

exports.register = function (server, options, next) {

  server.route({
    method: 'POST',
    path: '/webhook-endpoint',
    config: {
      description: 'Integration',
      tags: ['api'],
      validate: {
        options: {
          allowUnknown: true,
          stripUnknown: true
        },
        payload: Joi.object().options({abortEarly: false}).keys({
          event: Joi.string().required(),
          date: Joi.string().required(),
          payload: Joi.object().required().keys({
            id: Joi.string()
          }),
          user: Joi.object().required().keys({
            firstName: Joi.string().required(),
            lastName: Joi.string().required(),
            email: Joi.string().required()
          })
        })
      }
    },
    handler: function(request, reply) {
      console.log('You hooked me up! Event: ' + request.payload.event);
      reply();
    }
  });

  next();
};

exports.register.attributes = {
  pkg: require('../package.json')
};

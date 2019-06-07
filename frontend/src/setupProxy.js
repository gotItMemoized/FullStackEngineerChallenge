const proxy = require('http-proxy-middleware');

// setting up this so the docker container can have http://backend:8000
// but running npm start by itself will point to local host

// https://facebook.github.io/create-react-app/docs/proxying-api-requests-in-development#configuring-the-proxy-manually
module.exports = function(app) {
  var proxyTarget = 'http://localhost:8000/';
  if (process.env.DEV_FRONTEND_PROXY) {
    proxyTarget = process.env.DEV_FRONTEND_PROXY;
  }
  app.use(
    proxy('/api', {
      target: proxyTarget,
      pathRewrite: {
        '^/api/': '/', // remove base path
      },
    }),
  );
};

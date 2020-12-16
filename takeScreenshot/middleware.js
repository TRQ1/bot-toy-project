import routes from './routes.js';

export const Middleware = (req, res, next) => {
    res.locals.routes = routes;
    next()
};
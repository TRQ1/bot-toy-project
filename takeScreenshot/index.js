'use strict';

import express from 'express';
import { Middleware } from './middleware.js';
import routes from './routes.js';
import screenshotRouter from './routers/screenshotRouter.js';

const app = express();

app.use(Middleware);
app.use('/', screenshotRouter);
app.use(routes.screenshot, screenshotRouter);

export default app;

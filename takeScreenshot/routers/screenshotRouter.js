import express from 'express';
import routes from '../routes.js'
import { getMain, postScreenshot } from '../controllers/screenshotController.js';

const globalRouter = express.Router();

globalRouter.use(express.urlencoded({extended: true}));
globalRouter.use(express.json());

/**
 * Global redirect list
 */
globalRouter.get('/', getMain);
globalRouter.post(routes.screenshot, postScreenshot);

export default globalRouter;
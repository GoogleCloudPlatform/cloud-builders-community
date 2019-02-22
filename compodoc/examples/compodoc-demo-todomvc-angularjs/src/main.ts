import * as angular from 'angular';

//const uiRouter = require('angular-ui-router');

import uiRouter from 'angular-ui-router';

import { AppModule } from './app/';

angular.bootstrap(document, [
    AppModule,
    uiRouter
]);

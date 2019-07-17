import { AboutComponent } from './about.component';

export const AboutModule = angular
    .module('app.about', [])
    .component('about', AboutComponent)
    .name;

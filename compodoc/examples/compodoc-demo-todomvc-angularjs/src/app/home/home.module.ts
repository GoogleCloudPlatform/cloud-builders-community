import { HomeComponent } from './home.component';

export const HomeModule = angular
    .module('app.home', [])
    .component('home', HomeComponent)
    .name;

import { FooterComponent } from './footer.component';

export const FooterModule = angular
    .module('app.footer', [])
    .component('foot', FooterComponent)
    .name;

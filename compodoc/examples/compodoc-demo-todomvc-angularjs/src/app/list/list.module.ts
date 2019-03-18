import { ListComponent } from './list.component';
import { TodoModule } from './todo';

export const ListModule = angular
    .module('app.list', [
        TodoModule
    ])
    .component('list', ListComponent)
    .name;

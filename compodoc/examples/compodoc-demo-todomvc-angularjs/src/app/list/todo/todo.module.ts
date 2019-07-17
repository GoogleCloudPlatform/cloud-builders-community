import { TodoComponent } from './todo.component';

export const TodoModule = angular
    .module('app.todo', [])
    .component('todo', TodoComponent)
    .name;

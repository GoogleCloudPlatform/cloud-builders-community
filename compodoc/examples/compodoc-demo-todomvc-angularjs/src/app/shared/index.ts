import { TodoStore } from './services/todo.store';
import { EmitterService } from './services/emitter.service';
import { DoNothingDirective } from './directives/do-nothing.directive';

export const SharedModule = angular
    .module('app.shared', [])
    .service('TodoStore', TodoStore)
    .service('EmitterService', EmitterService)
    .directive('doNothing', () => new DoNothingDirective)
    .name;

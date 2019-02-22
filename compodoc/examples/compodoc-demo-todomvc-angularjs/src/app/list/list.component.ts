import { TodoStore } from '../shared/services/todo.store';
import { EmitterService } from '../shared/services/emitter.service';
import { Todo } from '../shared/models/todo.model';

export const ListComponent = {
    template: require('./list.component.html'),
    controller: class ListComponent {
        /**
         * Local reference of TodoStore
         */
        todoStore: TodoStore;
        EmitterService: EmitterService;
        id: string = 'ListComponent';
        todos: Array<Todo>;

        constructor(TodoStore: TodoStore, EmitterService: EmitterService) {
            'ngInject';
            let that = this;
    		this.todoStore = TodoStore;
            this.EmitterService = EmitterService;
            this.todos = TodoStore.getAll();

            this.EmitterService.attach('removeCompleted', this.id, function(currentFilter) {
                switch (currentFilter) {
                    case 'all':
                        that.todos = that.todoStore.getAll();
                        break;
                    case 'remaining':
                        that.todos = that.todoStore.getRemaining();
                        break;
                    case 'completed':
                        that.todos = that.todoStore.getCompleted();
                        break;
                }
            });

            this.EmitterService.attach('displayRemaining', this.id, function() {
                that.todos = that.todoStore.getRemaining();
            });
            this.EmitterService.attach('displayCompleted', this.id, function() {
                that.todos = that.todoStore.getCompleted();
            });
            this.EmitterService.attach('displayAll', this.id, function() {
                that.todos = that.todoStore.getAll();
            });
    	}
    }
};

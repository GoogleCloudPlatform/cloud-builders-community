import { TodoStore } from '../shared/services/todo.store';
import { EmitterService } from '../shared/services/emitter.service';

export const FooterComponent = {
    template: require('./footer.component.html'),
    controller: class FooterComponent {
        /**
         * Local reference of TodoStore
         */
        todoStore: TodoStore;
        EmitterService: EmitterService;
        currentFilter: string = 'all';

        constructor(TodoStore: TodoStore, EmitterService: EmitterService) {
            'ngInject';
    		this.todoStore = TodoStore;
            this.EmitterService = EmitterService;
    	}

        removeCompleted() {
    		this.todoStore.removeCompleted();
            this.EmitterService.notify('removeCompleted', this.currentFilter);
    	}

        displayCompleted() {
            this.currentFilter = 'completed';
            this.EmitterService.notify('displayCompleted');
        }

        displayRemaining() {
            this.currentFilter = 'remaining';
            this.EmitterService.notify('displayRemaining');
        }

        displayAll() {
            this.currentFilter = 'all';
            this.EmitterService.notify('displayAll');
        }
    }
};

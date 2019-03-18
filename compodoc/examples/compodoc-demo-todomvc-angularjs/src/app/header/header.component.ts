import { TodoStore } from '../shared/services/todo.store';

export const HeaderComponent = {
    template: require('./header.component.html'),
    controller: class HeaderComponent {
        /**
         * Application main title
         */
        title: string = 'todos';

        /**
         * Local reference of TodoStore
         */
        todoStore: TodoStore;

        newTodoText: string = '';

        constructor(TodoStore: TodoStore) {
            'ngInject';
    		this.todoStore = TodoStore;
    	}

        addTodo(event: any) {
            if (event.keyCode !== 13) return;
            if (this.newTodoText.trim().length) {
                this.todoStore.add(this.newTodoText);
    			this.newTodoText = '';
    		}
    	}
    }
};

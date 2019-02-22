import { Todo } from '../../shared/models/todo.model';
import { TodoStore } from '../../shared/services/todo.store';

export const TodoComponent = {
    template: require('./todo.component.html'),
    bindings: {
        todo: '<'
    },
    controller: class HeaderComponent {

        todo: Todo;
        editedTodo: Todo;
        todoStore: TodoStore;

        constructor(TodoStore: TodoStore) {
    		this.todoStore = TodoStore;
    	}

        remove(todo: Todo){
    		this.todoStore.remove(todo);
    	}

        toggleCompletion(todo: Todo) {
            todo.completed = !todo.completed;
    		this.todoStore.toggleCompletion(todo);
    	}

        editTodo(todo: Todo) {
            todo.editing = true;
            this.editedTodo = todo;
    	}

        stopEditing(todo: Todo, editedTitle: string) {
    		todo.title = editedTitle;
    		todo.editing = false;
    	}

    	cancelEditingTodo(todo: Todo) {
    		todo.editing = false;
    	}

    	updateEditingTodo(todo: Todo, editedTitle: string) {
    		editedTitle = editedTitle.trim();
    		todo.editing = false;
            this.editedTodo = null;

    		if (editedTitle.length === 0) {
    			return this.todoStore.remove(todo);
    		}

    		todo.title = editedTitle;

            this.todoStore.update();
    	}
    }
}

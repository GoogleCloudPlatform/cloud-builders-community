/**
 * The todo class
 */
export class Todo {
    /**
     * Completed status
     */
    completed: Boolean;
    /**
     * Editing status
     */
    editing: Boolean;

    /**
     * Title
     */
    private _title: String;
    get title() {
        return this._title;
    }
    set title(value: String) {
        this._title = value.trim();
    }

    static classMethod() {
        return 'hello';
    }

    constructor(title: String) {
        this.completed = false;
        this.editing = false;
        this.title = title.trim();
    }

    /**
     *  fakeMethod !!
     */
    fakeMethod(): boolean {
        return true;
    }
}

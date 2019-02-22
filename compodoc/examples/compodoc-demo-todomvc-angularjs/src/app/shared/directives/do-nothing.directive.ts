/**
 * This directive does nothing !
 */
export class DoNothingDirective implements angular.IDirective {
    restrict: string;

    constructor() {
        this.restrict = 'A';
    }

    link() {
        console.log('Do nothing directive');
    }
}

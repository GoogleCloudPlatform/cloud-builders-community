export function AppRouting($stateProvider: ng.ui.IStateProvider, $urlRouterProvider: ng.ui.IUrlRouterProvider) {
    let homeState = {
            name: 'home',
            url: '/',
            component: 'home'
        },

        aboutState = {
            name: 'about',
            url: '/about',
            component: 'about'
        };

    $stateProvider.state(homeState);
    $stateProvider.state(aboutState);
    $urlRouterProvider.otherwise('/');
}

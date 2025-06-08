import router from '@/router';

function goRoute(route: {}) {
    router.push(route)
}

export {
    goRoute
}
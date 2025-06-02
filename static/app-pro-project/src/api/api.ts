

function fetchSideMenu(): Promise<{}> {
    return fetch("/api/v1/get/menu/side", {})
        .then(response => response.json());
}

export { 
    fetchSideMenu
};
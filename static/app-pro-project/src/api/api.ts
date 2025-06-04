import type { ISideMenuItem } from "./api_interfaces"

const baseUrl = "http://localhost:3000"

/**
 * Класс ошибки запроса.
 */
class ErrorFetch extends Error {
    status: number;

    constructor(public code: number, public message: string) {
        super(message)
        this.status = code
    }
}

/**
 * Запрос меню приложения.
 * @returns json состава меню приложения.
 */
function fetchAppMenu(): Promise<ISideMenuItem[]> {
    return fetch(baseUrl + "/api/v1/get/menu/app", { })
        .then(response => {
            if (!response.ok) throw new ErrorFetch(response.status, response.statusText)

            return response.json()
        })
}

export { 
    fetchAppMenu,
    ErrorFetch
};


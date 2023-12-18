//https://eckertalex.dev/blog/typescript-fetch-wrapper
//https://github.com/EHB-TI/programming-project-groep-3_brussel-student-guide/blob/main/frontend/src/fetch.ts

export async function http<T>(path: string, config: RequestInit): Promise<T> {
    const request = new Request(path, config);
    const response = await fetch(request);

    if (response.status === 401) {
        console.error('Unauthorized access detected');
        throw new Error('Unauthorized');
    }

    if (!response.ok) {
        throw new Error('Error fetching.');
    }

    try {
        const json = await response.json();
        const object = json.result;
        if (object == null) {
            throw new Error(json.error);
        }
        return object;
    } catch (error) {
        console.error('Error parsing JSON response:', error);
        throw new Error('Error parsing JSON response');
    }
}

export async function rawhttp<T>(request: Request): Promise<T> {
    const response = await fetch(request);

    if (!response.ok) {
        throw new Error('Error fetching.');
    }

    try {
        const json = await response.json();
        const object = json.result;
        if (object == null) {
            throw new Error(json.error);
        }
        return object;
    } catch (error) {
        console.error('Error parsing JSON response:', error);
        throw new Error('Error parsing JSON response');
    }
}

export async function loginpost<T>(url: string, body: Object): Promise<T> {
    const body_string = JSON.stringify(body);
    const length = new TextEncoder().encode(body_string).length;
    const request = new Request(url, {
        method: 'POST',
        body: body_string,
        headers: {
            'Content-Type': 'application/json; charset=UTF-8',
            'Content-Length': length.toString(),
        },
    });

    try {
        const response = await fetch(request);

        if (!response.ok) {
            const json = await response.json();
            throw json.error;
        }

        const json = await response.json();
        const object = json.result;
        return object;
    } catch (error) {
        console.error('Internal error, try again later:', error);
        throw new Error('Internal error, try again later.');
    }
}

export async function get<T>(path: string, config?: RequestInit): Promise<T> {
    const init = { method: 'get', ...config };
    return await http<T>(path, init);
}

export async function delete_<T>(path: string, config?: RequestInit): Promise<T> {
    const init = { method: 'delete', ...config };
    return await http<T>(path, init);
}

export async function post<T, U>(path: string, body: T, config?: RequestInit): Promise<U> {
    const init = { method: 'post', body: JSON.stringify(body), ...config };
    return await http<U>(path, init);
}

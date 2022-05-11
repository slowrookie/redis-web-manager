export interface RequestOptions {
    method: string
    params?: Array<any> | {}
}

export class APIService {

    request = async (opt: RequestOptions): Promise<any> => {
        const go = (window as any).go;
        const method = go.main.App[opt.method];
        if (!method) return new Promise((resolve, reject) => {
            const msg = `Mehthod ${opt.method} not found`;
            return reject(msg)
        })

        console.log("Request: ", opt);
        return ((opt.params ? method(opt.params) : method()) as Promise<any>).then(v => {
            console.log("Response: ", v)
            return v;
        });
    }

}

<<<<<<< HEAD
export const defaultService = new APIService();
=======
export const isWails = !!(window as any).go;

export const defaultService = isWails ? defaultWailsService : defaultWebSocketService;
>>>>>>> db8c865 (feat: auto suggestion command)

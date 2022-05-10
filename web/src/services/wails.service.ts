import { APIService } from "./api.service";

export interface RequestOptions {
    method: string
    params?: Array<any> | {}
}

export class WailsService implements APIService {

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

export const defaultWailsService = new WailsService();
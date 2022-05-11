import { defaultWailsService } from "./wails.service";
import { defaultWebSocketService } from "./websocket.service";

export interface RequestOptions {
    method: string
    params?: Array<any> | {}
}

export interface APIService {

 request: (opt: RequestOptions) => Promise<any>

}

export const isWails = !!(window as any).go;

export const defaultService = isWails ? defaultWailsService : defaultWebSocketService;
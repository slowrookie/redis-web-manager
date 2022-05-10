import { defaultWailsService } from "./wails.service";
import { defaultWebSocketService } from "./websocket.service";

export interface RequestOptions {
    method: string
    params?: Array<any> | {}
}

export interface APIService {

 request: (opt: RequestOptions) => Promise<any>

}

export const defaultService = (window as any).go ? defaultWailsService : defaultWebSocketService;
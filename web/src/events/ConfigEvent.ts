import { Subject } from 'rxjs';

// change config event
export enum ConfigEventAction {
    Language,
    Theme,
    Port
}
export interface IConfigEvent {
    action: ConfigEventAction,
    params?: any
}
export const ConfigEvent = new Subject<IConfigEvent>();
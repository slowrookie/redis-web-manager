import { defaultService } from "./api.service";

export interface Lua {
  connectionID: string
  ID: string
  Desc: string
  Keys: Array<string>
  Args: Array<any>
  Script: string
  LastExecutionAt: number
  Elapsed: string
}

export const eidtLua = (lua: Lua): Promise<any> => {
  return defaultService.request({ method: lua.ID ? 'EditLua' : 'NewLua', params: lua });
}

export const deleteLua = (lua: Lua): Promise<any> => {
  return defaultService.request({ method: 'DeleteLua', params: lua });
}

export const loadLuas = (connectionId: string): Promise<Lua[]> => {
  return defaultService.request({ method: 'LoadLuas', params: connectionId });
}

export const executionScript = <T>(lua: Lua): Promise<T> => {
  return defaultService.request({ method: 'ExecutionScript', params: lua });
}
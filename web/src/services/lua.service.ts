import { defaultService } from "./api.service";

export interface Lua {
  connectionID: string
  id: string
  name: string
  keys: Array<string>
  args: Array<any>
  script: string
  lastExecutionAt: number
  elapsed: string
  result: any
}

export const eidtLua = (lua: Lua): Promise<any> => {
  return defaultService.request({ method: lua.id ? 'EditLua' : 'NewLua', params: lua });
}

export const deleteLua = (lua: Lua): Promise<any> => {
  return defaultService.request({ method: 'DeleteLua', params: lua });
}

export const loadLuas = (connectionId: string): Promise<Lua[]> => {
  return defaultService.request({ method: 'LoadLuas', params: connectionId });
}

export const executionScript = (lua: Lua): Promise<Lua> => {
  return defaultService.request({ method: 'ExecutionScript', params: lua });
}
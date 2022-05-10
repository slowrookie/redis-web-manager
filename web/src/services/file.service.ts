import { defaultService } from "./api.service";

export const readFile = (): Promise<string> => {
  return defaultService.request({method: 'ReadFile'})
}
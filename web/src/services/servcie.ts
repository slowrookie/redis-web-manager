export const defaultHeaders: HeadersInit = {
  'Content-Type': 'application/json'
}

export const serverErrorHandle = (response: Response): Promise<any> => {
  return response.text()
    .then(body => {
      try {
        return JSON.parse(body);
      } catch {
        throw Error(body);
      }
    });
}
export const CONNECTIONS = "CONNECTIONS";

export const KeyTypes = {
  STRING: 'STRING',
  LIST: 'LIST',
  SET: 'SET',
  ZSET: 'ZSET',
  HASH: 'HASH',
  STREAM: 'STREAM'
}

export const ComponentType = {
  INFO: 'INFO',
  CONSOLE: 'CONSOLE',
  STRING: 'STRING',
  LIST: 'LIST',
  SET: 'SET',
  HASH: 'HASH',
  ZSET: 'ZSET',
  STREAM: 'STREAM'
}

export const parseInfo = (res: string) => {
  const infoObject: any = {};
  const categories = res.split("#");
  categories.forEach(c => {
    const lines = c.split("\r\n"), partObject: any = {};
    if (lines[0].trim()) {
      var key = lines[0].trim()
      lines.forEach((e, i) => {
        const parts = lines[i].split(":");
        if (parts[1]) {
          partObject[parts[0]] = parts[1];
        }
      });
      infoObject[key] = partObject;
    }
  });
  return infoObject;
}

export const parseClient = (res: string) => {
  const lines = res.split("\n"), clients: Array<any> = [];
  lines.forEach(c => {
    var client: any = {}
    c && c.split(" ").forEach(v => {
      var kv = v.split("=");
      if (kv && kv.length === 2) {
        client[kv[0]] = kv[1];
      }
    })
    Object.keys(client).length && clients.push(client)
  })
  return clients;
}

export const parseArrayToObject = (arr: any) => {
  if (!Array.isArray(arr)) return {};
  const obj: any = {};
  [...Array(arr.length / 2)].forEach((_, i) => {
    const index = 2 * i;
    obj[arr[index]] = Array.isArray(arr[index + 1]) ? parseArrayToObject(arr[index + 1]) : arr[index + 1];
  })
  return obj;
}
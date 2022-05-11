import { filter, first, lastValueFrom, ReplaySubject, Subject, retry } from 'rxjs';
import { WebSocketSubject } from 'rxjs/webSocket';
import { v1 as uuidv1 } from 'uuid';

const JSONRPC_VERSION = '2.0'

export interface IJSONRPCRequest {
  jsonrpc?: string
  method: string
  params?: Array<any> | {}
  id?: string | number
}

export interface IJSONRPCResponse {
  id?: string
  result: any
  error: any
}

export enum ConnectionState {
  CONNECTED = "Connected",
  CONNECTING = "Connecting",
  CLOSING = "Closing",
  DISCONNECTED = "Disconnected"
}

export class WebSocketService {
  private connectionState = new ReplaySubject<ConnectionState>(1);
  private socket: WebSocketSubject<ArrayBuffer | Object>;

  private messageObserver = new Subject<any>();
  private binaryObserver = new Subject<ArrayBuffer>();

  constructor(url: string) {
    this.connectionState.next(ConnectionState.CONNECTING);

    this.socket = new WebSocketSubject({
      binaryType: 'arraybuffer',
      url,
      openObserver: {
        next: () => this.connectionState.next(ConnectionState.CONNECTED)
      },
      closingObserver: {
        next: () => this.connectionState.next(ConnectionState.CLOSING)
      },
      closeObserver: {
        next: () => this.connectionState.next(ConnectionState.DISCONNECTED)
      },
      deserializer: (e: MessageEvent) => {
        try {
          if (e.data instanceof ArrayBuffer) {
            return e.data;
          } else {
            return JSON.parse(e.data);
          }
        } catch (e) {
          console.error(e);
          return null;
        }
      }
    })

    // message
    this.socket.pipe(
      retry(),
      filter((v: any, index: number) => !(v instanceof ArrayBuffer))
    ).subscribe(message => {
      this.messageObserver.next(message);
    })

    // binary message
    this.socket.pipe(
      retry(),
      filter((value: any, index: number) => value instanceof ArrayBuffer)
    ).subscribe(message => {
      this.binaryObserver.next(message);
    })

    this.connectionState.subscribe((state) => {
      console.log(`WebSocket state ${state}`);
    });

  }

  request = async (request: IJSONRPCRequest): Promise<any> => {
    if (!request.jsonrpc) request.jsonrpc = JSONRPC_VERSION;
    if (!request.id) request.id = uuidv1();
    if (!request.params) request.params = [];

    this.socket.next(request);

    const obs = this.messageObserver.pipe(
      filter((v: any) => request.id === v.id),
      first()
    );

    return lastValueFrom(obs).then((message: IJSONRPCResponse) => {
      if (message.error) {
        throw message.error.message || message.error;
      }
      return message.result;
    });
  }

}

let address = process.env.NODE_ENV === 'development' ? 'ws://localhost:63790/ws' : `ws://${window.location.host}/ws`;

export const defaultWebSocket = new WebSocketService(address);
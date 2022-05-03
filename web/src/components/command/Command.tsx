import React, { useEffect, useRef, useState } from 'react';
import { debounceTime, fromEvent, Subject } from 'rxjs';
import { ITerminalOptions, Terminal } from 'xterm';
import { FitAddon } from 'xterm-addon-fit';
import 'xterm/css/xterm.css';
import { Connection, executeCommand } from '../../services/connection.service';

const formatArray = (arr: Array<any>, index: number, tab: number): any => {
  return arr.map((v, i) => {
    if (Array.isArray(v)) {
      return formatArray(v, i + 1, tab + 1)
    }
    var prefix = '';
    if (index > 0) {
      prefix += (i === 0 ? `${index}) ` : [...Array(tab)].fill('   ').join())
    }
    return prefix + `${i + 1}) ${v}`;
  }).join('\r\n');
}

const formatRet = (ret: any) => {
  if (Array.isArray(ret)) {
    return formatArray(ret, 0, 0);
  }
  return ret;
}

const fitAddon = new FitAddon();

export interface ICommandProps extends ITerminalOptions {
  connection: Connection
}

export const Command = (props: ICommandProps) => {
  // const theme = useTheme(), { t } = useTranslation();

  const container = useRef<HTMLDivElement>(null),
    [xterm, setXterm] = useState<Terminal>(),
    [command, setCommand] = useState(''),
    [selectedDB, setSelectedDB] = useState<number>(0),
    [executeEvent] = useState(new Subject<string>());

  useEffect(() => {
    const sub = executeEvent.subscribe(v => {
      if (!command || !xterm) return;
    
      let commands: Array<Array<any>> = [[]];
      let currentDB: number = selectedDB;
      let commandSplit = command.trim().split(' ');
      let firstToken = commandSplit[0].toUpperCase();
     
      switch(firstToken) {
        case 'SELECT':
          if (commandSplit.length > 1) {
            currentDB = Number(commandSplit[1]);
          }
          commands = [commandSplit];
          break;
        case 'CLEAR':
          setCommand('');
          xterm.clear();
          xterm.write('\x1bc')
          xterm.write(`\r\n${currentDB} > `);
          return;
        default:
          commands = [['SELECT', currentDB], commandSplit];
          break;
      }
  
      executeCommand<Array<any>>({ id: props.connection.id, commands })
        .then((ret) => {
          xterm.write(`\r\n`);
          xterm.writeln(formatRet(ret))
          if (commandSplit[0].toUpperCase() === 'SELECT') {
            setSelectedDB(currentDB);
          }
          xterm.write(`\r\n${currentDB} > `);
        })
        .catch((err: Error) => { 
          xterm.write(`\r\n`);
          xterm.writeln(String(err));
          xterm.write(`\r\n${selectedDB} > `);
        })
        .finally(() => {
          setCommand('');
        });
    });

    return () => sub && sub.unsubscribe();
  }, [props.connection.id, executeEvent, command, xterm, selectedDB])

  useEffect(() => {
    if (!xterm) return;

    xterm.onData(e => {
      switch (e) {
        case '\u0003': // Ctrl+C
        xterm.write('^C');
          break;
        case '\r': // Enter
          executeEvent.next('');
          break;
        case '\u007F': // Backspace (DEL)
          // Do not delete the prompt
          if (((xterm) as any)._core.buffer.x > 4) {
            xterm.write('\b \b');
            setCommand(v => {
              if (v.length > 0) {
                return v.substring(0, v.length - 1)
              }
              return v;
            })
          }
          break;
        default: // Print all other characters
          if ((e >= String.fromCharCode(0x20) && e <= String.fromCharCode(0x7E)) || e >= '\u00a0') {
            setCommand(v => {
              v += e;
              return v;
            })
            xterm.write(e);
          }
      }
    });

    return () => {
      xterm.onData(() => {})
    }
  }, [xterm, executeEvent])

  // init xterm
  useEffect(() => {
    if (!container || !container.current) return;

    let term = new Terminal({...props, cursorBlink: true});
    
    term.loadAddon(fitAddon);
    term.open(container.current);

    term.write(`\r\n0 > `);

    fitAddon.fit();
    setXterm(term)

    return () => {
      term.dispose();
      setXterm(undefined);
    }
  }, [container, props])

  // resize & fit
  useEffect(() => {
    const sub = fromEvent(window, "resize").pipe(
      debounceTime(250)
    ).subscribe(v => {
      fitAddon.fit();
    })
    return () => sub && sub.unsubscribe();
  })

  return (
    <>
      <div style={{height: '100%'}} ref={container}></div>
    </>
  )
}
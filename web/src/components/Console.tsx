import { ContextualMenu, DirectionalHint, IContextualMenuItem, Stack, Target, useTheme, TextField } from '@fluentui/react';
import React, { ChangeEvent, CSSProperties, KeyboardEvent, useEffect, useRef, useState } from 'react';
import { useTranslation } from 'react-i18next';
import { Connection, executeCommand } from '../services/connection.service';
import { ErrorMessageBar } from './common/ErrorMessageBar';


export interface IConsoleProps {
  connection: Connection
}

declare type InputLine = 'I' | 'O'

interface IInputLine {
  type: InputLine
  ret: string
  db?: number
  command?: string
  disabled?: boolean
}

const defaultInputLine: IInputLine = { type: 'I', command: '', disabled: false, ret: '', db: 0 };

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

export const Console = (props: IConsoleProps) => {
  const theme = useTheme();
  const { t } = useTranslation();
  const inputStyle: CSSProperties = {
    border: 'none',
    outline: 'none',
    color: theme.palette.themePrimary,
    height: 14,
    lineHeight: '14px',
    width: '100%'
  }

  const [lines, setLines] = useState<Array<IInputLine>>([defaultInputLine]),
    [currentLine, setCurrentLine] = useState(''),
    [error, setError] = useState<Error>(),
    [selectedDB, setSelectedDB] = useState<number>(0),
    [showContextualMenu, setShowContextualMenu] = useState(false),
    [contextualMenuTarget, setContextualMenuTarget] = useState<Target>(),
    consoleDiv = useRef<HTMLDivElement | null>(null);

  const handleKeyPress = (e: KeyboardEvent) => {
    if (e.key === "Enter") {
      setError(undefined);
      var commands: Array<Array<any>> = [[]];
      var currentCommand = currentLine.trim().split(" ");
      var currentDB: number = selectedDB;
      // clear
      if (currentCommand[0].toUpperCase() === 'CLEAR') {
        setLines([{ ...defaultInputLine, db: currentDB }]);
        return;
      }
      // select
      if (currentCommand[0].toUpperCase() === 'SELECT') {
        if (currentCommand.length > 1) {
          currentDB = Number(currentCommand[1]);
          setSelectedDB(Number(currentCommand[1]));
        }
        commands = [currentCommand];
      } else {
        commands = [['SELECT', currentDB], currentCommand];
      }

      executeCommand<Array<any>>({ id: props.connection.id, commands })
        .then((ret) => {
          console.log(ret);
          lines[lines.length - 1] = { ...lines[lines.length - 1], disabled: true, command: currentLine }
          lines.push({ type: 'O', ret: formatRet(ret[ret.length - 1]) });
          lines.push({ ...defaultInputLine, db: currentDB });
          setLines([...lines]);
          setCurrentLine('');
        })
        .catch((err: Error) => { setError(err); });
    }
  }

  const handleChange = (e: ChangeEvent<any>) => {
    setCurrentLine(e.currentTarget.value);
  }

  const contextMenuItems: IContextualMenuItem[] = [
    {
      key: 'clear',
      text: t('Clear'),
      iconProps: { iconName: 'removeFromShoppingList' },
      onClick: () => {
        setLines([{ ...defaultInputLine, db: selectedDB }]);
      },
    }
  ]

  useEffect(() => {
    if (!consoleDiv || !consoleDiv.current) {
      return;
    }
    // contextmenu
    consoleDiv.current.addEventListener("contextmenu", (e: globalThis.MouseEvent) => {
      e.stopPropagation();
      e.preventDefault();
      setContextualMenuTarget(e);
      setShowContextualMenu(true);
    });
    return () => {
      document.removeEventListener("contextmenu", () => { });
    }
  }, [consoleDiv])

  return (
    <div style={{ height: "100%" }} ref={consoleDiv}>
      <Stack style={{ height: '100%' }}>

        <ErrorMessageBar error={error} />

        <Stack.Item grow={1} style={{ padding: 5, overflow: 'auto', color: theme.palette.neutralPrimary, fontSize: 12 }}>
          {lines && lines.map((line, i) => <Stack key={i} horizontal verticalAlign="center" tokens={{ childrenGap: 10 }}>

            {line.type === "I" && (<>
              <span style={{ color: theme.palette.themePrimary }}>{`${props.connection.name} : ${line.db} >`}</span>
              <Stack.Item grow={1}>
                {line.disabled ? (<span style={inputStyle}>{line.command}</span>) :
                  (
                    <TextField styles={{root: {height: 24}, fieldGroup: {height: 24}}} autoFocus underlined onKeyPress={handleKeyPress} onChange={handleChange} />
                  // <input style={inputStyle} autoFocus onKeyPress={handleKeyPress} onChange={handleChange}></input>
                  )}
              </Stack.Item>
            </>)}

            {line.type === "O" && <div><pre style={{ margin: 5 }}>{line.ret}</pre></div>}

          </Stack>)}
        </Stack.Item>
      </Stack>

      <ContextualMenu
        directionalHint={DirectionalHint.rightCenter}
        items={contextMenuItems}
        hidden={!showContextualMenu}
        target={contextualMenuTarget}
        onDismiss={() => setShowContextualMenu(false)}
      />
    </div>
  )
}
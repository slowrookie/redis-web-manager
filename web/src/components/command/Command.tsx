import { BasePicker, Callout, FocusZone, FocusZoneDirection, List, SearchBox, Stack, useTheme } from '@fluentui/react';
import { useBoolean, useId } from '@fluentui/react-hooks';
import React, { ChangeEvent, CSSProperties, useState } from 'react';
import { useTranslation } from 'react-i18next';
import { Connection, executeCommand, suggestions } from '../../services/connection.service';
import { ErrorMessageBar } from '../common/ErrorMessageBar';
import { Suggestion } from './Suggestion';


export interface ICommandProps {
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
      prefix += (i === 0 ? `${index}) ` : [...Array(tab)].fill('\t').join())
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

export const Command = (props: ICommandProps) => {
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
    searchBoxId = useId(`command-callout-button-${props.connection.id}`),
    [isCalloutVisible, { toggle: toggleIsCalloutVisible, setFalse: setCalloutVisibleFalse, setTrue: setCalloutVisibleTrue }] = useBoolean(false),
    [expects, setExpects] = useState<Array<any>>([]);

  const handleSearch = (v?: string) => {
    if (!v) return;

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
        lines[lines.length - 1] = { ...lines[lines.length - 1], disabled: true, command: currentLine }
        lines.push({ type: 'O', ret: formatRet(ret[ret.length - 1]) });
        lines.push({ ...defaultInputLine, db: currentDB });
        setLines([...lines]);
        setCurrentLine('');
      })
      .catch((err: Error) => { setError(err); });
  }

  const handleChange = (e?: ChangeEvent<any>, value?: string) => {
    setCurrentLine(value || "");
    if (value) {
      suggestions(value.trim()).then((v: any) => {
        console.log(v);
        setExpects(v);
        v && v.length && setCalloutVisibleTrue();
      })
    }
  }

  return (
    <div style={{ height: "100%" }}>
      <Stack style={{ height: '100%' }}>

        <ErrorMessageBar error={error} />

        <Stack.Item grow={1} style={{ padding: 5, overflow: 'auto', color: theme.palette.neutralPrimary, fontSize: 12 }}>
          {lines && lines.map((line, i) => <Stack key={i} horizontal verticalAlign="center" tokens={{ childrenGap: 10 }}>

            {line.type === "I" && (<>
              <span style={{ color: theme.palette.themePrimary }}>{`${line.db} >`}</span>
              <Stack.Item grow={1}>
                {line.disabled ? (<span style={inputStyle}>{line.command}</span>) :
                  (
                    <>
                      <Suggestion />
                      {/* <SearchBox
                        id={searchBoxId}
                        styles={{ root: { height: 24, paddingLeft: 0 }, iconContainer: { width: 0 } }}
                        autoFocus
                        iconProps={{ iconName: '' }}
                        underlined={true}
                        onChange={handleChange}
                        onSearch={handleSearch}
                      /> */}
                    </>
                  )}
              </Stack.Item>
            </>)}

            {line.type === "O" && <div style={{ width: 'calc(100vh -50px)', overflow: 'auto' }}>
              <pre style={{ margin: 5, whiteSpace: 'pre-wrap', wordWrap: 'break-word' }}>{line.ret}</pre>
            </div>}

          </Stack>)}
        </Stack.Item>
      </Stack>

      {isCalloutVisible && <>
        <Callout role="suggestion"
          gapSpace={2}
          target={`#${searchBoxId}`}
          onDismiss={toggleIsCalloutVisible}
          calloutMaxHeight={300}
          coverTarget={false}
          alignTargetEdge={true}
          setInitialFocus={true}
          isBeakVisible={false}
        >
          <FocusZone direction={FocusZoneDirection.vertical}>
            <List id='SearchList' tabIndex={0} items={expects} onRenderCell={(item: any) => {
              return (<div key={item} data-is-focusable={true}>
                {item}
              </div>)
            }}
            />
          </FocusZone>

        </Callout>
      </>}
    </div>
  )
}
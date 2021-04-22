import { DefaultButton, Stack, TooltipHost } from '@fluentui/react';
import React, { useCallback, useEffect, useState } from 'react';
import { useTranslation } from 'react-i18next';
import { Connection, executeCommand } from '../services/connection.service';
import { ErrorMessageBar } from './common/ErrorMessageBar';
import { FormatTextField } from './common/FormatTextField';
import { Loading } from './common/Loading';
import { KeyHeader } from './KeyHeader';

const buttonStyles = {
  root: {
    minWidth: 'auto',
    padding: '0 6px'
  }
}

export interface IStringKeyyProps {
  connection: Connection
  db: number
  component: string,
  keyName: string,
  onKeyNameChanged: (oldKeyName: string, keyName: string) => void
  onDeletedKey: (keyName: string) => void
}

export interface IStringKey {
  keyName: string
  TTL: number
  value: string
  initialValue: string
  length: number
}

const defaultStringKey: IStringKey = {
  keyName: '',
  TTL: -1,
  value: '',
  initialValue: '',
  length: 0
}

export const StringKey = (props: IStringKeyyProps) => {

  const { keyName, connection, db, onKeyNameChanged } = props;

  const { t } = useTranslation();

  const [keyProps, setKeyProps] = useState({ ...defaultStringKey, keyName: keyName }),
    [error, setError] = useState<string>(),
    [loading, setLoading] = useState(false);

  useEffect(() => {
    setKeyProps({ ...defaultStringKey, keyName })
  }, [keyName])

  const load = useCallback(() => {
    setLoading(true);
    setError('');
    executeCommand<Array<any>>({
      id: connection.id, commands: [
        ['SELECT', db],
        ['TTL', keyProps.keyName],
        ['GET', keyProps.keyName],
        ['STRLEN', keyProps.keyName]
      ]
    }).then((ret) => {
      if (!ret || !ret.length) return;
      console.log(ret);
      setKeyProps({
        keyName: keyProps.keyName,
        TTL: ret[1],
        value: ret[2],
        initialValue: ret[2],
        length: ret[3]
      });
    })
      .catch(err => setError(err))
      .finally(() => setLoading(false));

  }, [connection.id, db, keyProps.keyName]);

  useEffect(() => {
    load();
  }, [load]);

  const handleValue = () => {
    setError('');
    executeCommand({
      id: connection.id, commands: [
        ['SELECT', db],
        ['SET', keyProps.keyName, keyProps.value],
      ]
    }).then((ret) => {
      setKeyProps({ ...keyProps, initialValue: keyProps.value });
      load();
    })
      .catch(err => setError(err));
  }

  const textFieldActions = () => {
    return (
      <TooltipHost content={t('Update')}>
        <DefaultButton styles={buttonStyles} disabled={keyProps.initialValue === keyProps.value} iconProps={{ iconName: 'save', style: { height: 'auto' } }}
          onClick={handleValue}></DefaultButton>
      </TooltipHost>
    )
  }

  return (<>
    <Loading loading={loading}></Loading>
    <Stack style={{ height: '100%' }} tokens={{ padding: 5, childrenGap: 10 }}>
      {/* error */}
      <ErrorMessageBar error={error}></ErrorMessageBar>

      <KeyHeader {...props} keyName={keyProps.keyName} TTL={keyProps.TTL}
        onKeyNameChanged={(oldValue, newValue) => {
          setKeyProps({ ...keyProps, keyName: newValue });
          onKeyNameChanged(oldValue, newValue);
        }}
        onTTLChanged={(v) => { setKeyProps({ ...keyProps, TTL: v }) }}
        onRefresh={() => load()} />

      <Stack.Item grow={1}>
        <FormatTextField label={t('Value')} multiline value={keyProps.value} length={keyProps.length}
          onChange={(e, v) => v && setKeyProps({ ...keyProps, value: v })}
          actions={textFieldActions()}
        ></FormatTextField>
      </Stack.Item>

    </Stack >
  </>)
}
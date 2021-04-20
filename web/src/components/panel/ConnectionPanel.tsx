import { DefaultButton, MessageBar, MessageBarType, Overlay, Panel, PanelType, Pivot, PivotItem, PrimaryButton, Separator, Spinner, SpinnerSize, Stack, TextField } from '@fluentui/react';
import React, { useEffect, useState } from 'react';
import { useTranslation } from 'react-i18next';
import { Connection, saveConnection, testConnection } from '../../services/connection.service';

export interface IConnectionPanel {
  connection?: Connection
  isOpen: boolean,
  setIsOpen: (open: boolean) => void
  onSave?: (connection: Connection) => void
}

const defaultConnection: Connection = {
  id: '',
  name: '',
  host: '127.0.0.1',
  port: 6379,
  auth: '',
  username: '',
  keysPattern: '*',
  namespaceSeparator: ':',
  timeoutConnect: 20000,
  timeoutExecute: 60000,
  dbScanLimit: 20,
  dataScanLimit: 2000
}

export const ConnectionPanel = (props: IConnectionPanel) => {
  const { t } = useTranslation();
  const { connection, isOpen, setIsOpen, onSave } = props;

  const [_connection, _setConnection] = useState<Connection>(defaultConnection),
    [connecting, setConnecting] = useState(false),
    [error, setError] = useState<string>(),
    [success, setSuccess] = useState<string>()
    ;

  useEffect(() => {
    setError("");
    setSuccess("");
    connection && _setConnection(connection.id ? connection : defaultConnection);
  }, [connection])

  const onTestConnection = () => {
    setError("");
    setSuccess("");
    setConnecting(true);
    testConnection(_connection).then(() => {
      setSuccess(t("The test connection to the Redis server is successful!"))
    })
      .catch(err => setError(err))
      .finally(() => setConnecting(false))
  }

  const OnConnection = () => {
    setError("");
    setSuccess("");
    setConnecting(true);
    saveConnection(_connection).then((v) => {
      setSuccess(t("The connection to the Redis server is successful!"));
      setIsOpen(false);
      onSave && onSave(v);
    })
      .catch(err => setError(err))
      .finally(() => setConnecting(false));
  }

  return (
    <Panel
      isOpen={isOpen}
      isBlocking={true}
      type={PanelType.medium}
      isLightDismiss={true}
      onDismiss={() => setIsOpen(false)}
      headerText={t('Connection settings')}
      onRenderFooterContent={() => {
        const disabled = !(_connection.name && _connection.host && _connection.port);
        return (
          <Stack tokens={{ childrenGap: 10 }} horizontal horizontalAlign="space-evenly">
            <PrimaryButton
              disabled={disabled}
              text={t('Test Connection')}
              iconProps={{ iconName: 'plugConnected' }} onClick={onTestConnection}></PrimaryButton>
            <Stack.Item grow={1}><span></span></Stack.Item>
            <DefaultButton disabled={disabled} text={t('OK')} onClick={OnConnection}></DefaultButton>
            <DefaultButton text={t('Cancel')} onClick={() => setIsOpen(false)}></DefaultButton>
          </Stack>
        )
      }}
      isFooterAtBottom={true}
    >
      {connecting &&
        <Overlay styles={{ root: { zIndex: 1 } }}>
          <Stack style={{ height: '100%' }} horizontalAlign="center" verticalAlign="center">
            <Spinner size={SpinnerSize.large} label={t('Connecting...')} />
          </Stack>
        </Overlay>
      }

      <Stack tokens={{ childrenGap: 10, padding: 10 }}>
        <Pivot>
          <PivotItem headerText={t('Basic')}>
            <TextField label={t('Name')} placeholder={t('Connection name')} required value={_connection.name} onChange={(e, v) => {
              v && _setConnection(c => ({ ...c, name: v }));
            }} />
            <TextField label={t('Address')} placeholder={t('Service address')} required value={_connection.host} onChange={(e, v) => {
              v && _setConnection(c => ({ ...c, host: v }))
            }} />
            <TextField label={t('Port')} type="number" placeholder={t('Port')} min={0} max={65535} required value={`${_connection.port}`} onChange={(e, v) => {
              var nv = Number(v);
              if (nv > 65535) { nv = 65535 };
              _setConnection(c => ({ ...c, port: nv }));
            }} />
            <TextField label={t('Password')} type="password" placeholder={t('(Optional) Service authentication password')} canRevealPassword={true} value={_connection.auth} onChange={(e, v) => {
              v && _setConnection(c => ({ ...c, auth: v }))
            }}
            />
            <TextField label={t('Username')} placeholder={t('(Optional) Service authentication user name Reids> 6.0')} value={_connection.username} onChange={(e, v) => {
              v && _setConnection(c => ({ ...c, username: v }))
            }} />
          </PivotItem>

          <PivotItem headerText={t("Advanced")}>
            <Separator>{t('Key load')}</Separator>
            <TextField label={t('Default key filtering')} required value={_connection.keysPattern} onChange={(e, v) => {
              v && _setConnection(c => ({ ...c, keysPattern: v }))
            }} />
            <TextField label={t('Namespace separator')} required value={_connection.namespaceSeparator} onChange={(e, v) => {
              v && _setConnection(c => ({ ...c, namespaceSeparator: v }))
            }} />
            <Separator>{t('Timeouts and limits')}</Separator>
            <TextField label={t('Connection timeout (ms)')} type="number" placeholder={t('Connection timeout (ms)')} required value={`${_connection.timeoutConnect}`} onChange={(e, v) => {
              if (Number(v) < 0) return;
              _setConnection(c => ({ ...c, timeoutConnect: Number(v) }));
            }} />
            <TextField label={t('Execution timeout (ms)')} type="number" placeholder={t('Execution timeout (ms)')} required value={`${_connection.timeoutExecute}`} onChange={(e, v) => {
              if (Number(v) < 0) return;
              _setConnection(c => ({ ...c, timeoutExecute: Number(v) }));
            }} />
            <TextField label={t('Database display limit')} type="number" placeholder={t('Database display limit')} required value={`${_connection.dbScanLimit}`} onChange={(e, v) => {
              _setConnection(c => ({ ...c, dbScanLimit: Number(v) }))
            }} />
            <TextField label={t('Data scan limit')} type="number" placeholder={t('Data scan limit')} required value={`${_connection.dataScanLimit}`} onChange={(e, v) => {
              _setConnection(c => ({ ...c, dataScanLimit: Number(v) }))
            }} />
          </PivotItem>
        </Pivot>

        {error && <MessageBar styles={{ icon: { height: 16, lineHeight: '14px' } }} messageBarType={MessageBarType.blocked} isMultiline={true} onDismiss={() => { setError("") }} truncated={true}>
          {error}
        </MessageBar>}
        {success && <MessageBar styles={{ icon: { height: 16, lineHeight: '14px' } }} messageBarType={MessageBarType.success} isMultiline={false} onDismiss={() => { setSuccess("") }}>
          {success}
        </MessageBar>}

      </Stack>
    </Panel>
  )
}
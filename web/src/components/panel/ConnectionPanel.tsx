import { ChoiceGroup, CommandBarButton, DefaultButton, IChoiceGroupOption, IconButton, ITextFieldProps, Label, MessageBar, MessageBarType, Overlay, Panel, PanelType, Pivot, PivotItem, PrimaryButton, Separator, Spinner, SpinnerSize, Stack, TextField, Toggle } from '@fluentui/react';
import React, { useEffect, useRef, useState } from 'react';
import { useTranslation } from 'react-i18next';
import { Connection, saveConnection, testConnection } from '../../services/connection.service';
import { ErrorMessageBar } from '../common/ErrorMessageBar';

export interface IConnectionPanel {
  connection?: Connection
  isOpen: boolean,
  setIsOpen: (open: boolean) => void
  onSave?: (connection: Connection) => void
}

interface Addr {
  host: string,
  port: number
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
  timeoutConnect: 2000,
  timeoutExecute: 2000,
  dbScanLimit: 20,
  dataScanLimit: 2000,
  tls: {
    enable: false
  },
  ssh: {
    enable: false
  }
}

const ReadFileTextFiled = (props: ITextFieldProps) => {
  const { t } = useTranslation();
  const fileChosenRef = useRef<any>();

  return <TextField {...props} onRenderLabel={(fieldProps, defaultRender: any) => {
    return (
      <Stack horizontal>
        {defaultRender(fieldProps)}
        <Stack.Item grow={1}><span></span></Stack.Item>
        <input type="file" style={{ display: 'none' }} ref={fileChosenRef} onChange={(e) => {
          if (!e.target.files) return;
          let file = e.target.files[0];
          const fileReader = new FileReader()
          fileReader.onloadend = () => {
            props.onChange && props.onChange(e, fileReader.result?.toString());
          }
          fileReader.readAsText(file);
        }} />
        <CommandBarButton iconProps={{ iconName: 'Upload' }} text={t('Upload')} onClick={() => {
          fileChosenRef.current.click();
        }}></CommandBarButton>
      </Stack>
    )
  }} />
}


export const ConnectionPanel = (props: IConnectionPanel) => {
  const { t } = useTranslation();
  const { connection, isOpen, setIsOpen, onSave } = props;

  const [_connection, _setConnection] = useState<Connection>(defaultConnection),
    [connecting, setConnecting] = useState(false),
    [error, setError] = useState<Error | undefined>(),
    [success, setSuccess] = useState<string>()
    ;

  const clientTypeOptions: IChoiceGroupOption[] = [
    { key: 'general', text: t('general') },
    { key: 'cluster', text: t('cluster') },
    { key: 'sentinel', text: t('sentinel') }
  ];

  useEffect(() => {
    setError(undefined);
    setSuccess("");
    if (!connection || !connection.id) {
      _setConnection(defaultConnection)
    } else {
      _setConnection(connection)
    }
    connection && connection.addrs && setAddrs(connection.addrs.map(a => ({ host: a.split(':')[0], port: Number(a.split(':')[1] || 0) })));
  }, [connection])

  const onTestConnection = () => {
    setError(undefined);
    setSuccess("");
    setConnecting(true);
    testConnection(_connection).then(() => {
      setSuccess(t("The test connection to the Redis server is successful!"))
    })
      .catch((err: Error) => {
        setError(err)
      })
      .finally(() => setConnecting(false))
  }

  const OnConnection = () => {
    setError(undefined);
    setSuccess("");
    setConnecting(true);
    saveConnection(_connection).then((v) => {
      setSuccess(t("Save success!"));
      setIsOpen(false);
      onSave && onSave(v);
    })
      .catch(err => setError(err))
      .finally(() => setConnecting(false));
  }

  // address
  const [addrs, setAddrs] = useState<Array<Addr>>([{ host: '', port: 0 }]);
  useEffect(() => {
    _setConnection(c => ({ ...c, addrs: addrs.map(a => `${a.host}:${a.port}`) }))
  }, [addrs])

  const handleAddAddr = () => {
    setAddrs([...addrs, { host: '', port: 0 }])
  }

  const handleDeleteAddr = (index: number) => {
    addrs.splice(index, 1);
    setAddrs([...addrs]);
  }

  const general = (<>
    <TextField label={t('Address')} placeholder={t('Service address')} required value={_connection.host} onChange={(e, v) => {
      _setConnection(c => ({ ...c, host: v || '' }))
    }} />
    <TextField label={t('Port')} type="number" placeholder={t('Port')} min={0} max={65535} required value={`${_connection.port}`} onChange={(e, v) => {
      var nv = Number(v);
      if (nv > 65535) { nv = 65535 };
      _setConnection(c => ({ ...c, port: nv }));
    }} />
  </>);

  const cluster = (<>
    <Stack tokens={{ childrenGap: 10 }}>
      <Stack horizontal horizontalAlign='space-between' verticalAlign='center' tokens={{ childrenGap: 10 }}>
        <Label required>{t('Address')}</Label>
        <IconButton iconProps={{ iconName: 'circleAdditionSolid' }} onClick={handleAddAddr} />
      </Stack>
      <Stack tokens={{ childrenGap: 10 }}>
        {addrs.length && [...Array(addrs.length)].fill(0).map((_, index: number) => {
          return <Stack key={`address${index}`} horizontal tokens={{ childrenGap: 10 }} horizontalAlign='space-evenly'>
            <Stack.Item grow={1}>
              <TextField label={''} placeholder={t('Service address')} underlined value={addrs[index].host} onChange={(e, v) => {
                addrs[index].host = v || '';
                setAddrs([...addrs])
              }} />
            </Stack.Item>
            <TextField label={''} type="number" placeholder={t('Port')} underlined min={0} max={65535} value={`${addrs[index].port}`} onChange={(e, v) => {
              var nv = Number(v);
              if (nv > 65535) { nv = 65535 };
              addrs[index].port = nv || 0;
              setAddrs([...addrs])
            }} />
            <IconButton iconProps={{ iconName: 'delete' }} onClick={() => handleDeleteAddr(index)} />
          </Stack>
        })}
      </Stack>

      <ChoiceGroup label={t('Route')} required styles={{ flexContainer: { display: "flex", justifyContent: 'space-between' } }}
        selectedKey={_connection.routeByLatency ? 'RouteByLatency' : _connection.routeRandomly ? 'RouteRandomly' : 'RouteByLatency'}
        onChange={(e, o) => {
          if (!o) return;
          const routeByLatency = 'RouteByLatency' === o.key;
          const routeRandomly = 'RouteRandomly' === o.key;
          _setConnection(c => ({ ...c, routeByLatency, routeRandomly }))
        }}
        defaultSelectedKey="RouteByLatency" options={[
          { key: 'RouteByLatency', text: t('RouteByLatency') },
          { key: 'RouteRandomly', text: t('RouteRandomly') }
        ]} />

    </Stack>
  </>)

  const sentinel = (<>
    <TextField label={t('Sentinel password')} type="password" placeholder={t('Sentinel password')} canRevealPassword={true} value={_connection.sentinelPassword} onChange={(e, v) => {
      _setConnection(c => ({ ...c, sentinelPassword: v || '' }))
    }} />

    <TextField label={t('Master name')} type="text" placeholder={t('Master name')} value={_connection.masterName} onChange={(e, v) => {
      _setConnection(c => ({ ...c, masterName: v || '' }))
    }} />
  </>)

  const basic = (
    <PivotItem headerText={t('Basic')}>
      <TextField label={t('Name')} placeholder={t('Connection name')} required value={_connection.name} onChange={(e, v) => {
        _setConnection(c => ({ ...c, name: v || '' }));
      }} />
      <ChoiceGroup label={t('Type')} required styles={{ flexContainer: { display: "flex", justifyContent: 'space-between' } }}
        selectedKey={_connection.isCluster ? 'cluster' : _connection.isSentinel ? 'sentinel' : 'general'}
        onChange={(e, o) => {
          if (!o) return;
          _setConnection(c => ({ ...c, isCluster: o.key === 'cluster', isSentinel: o.key === 'sentinel' }))
        }}
        defaultSelectedKey="general" options={clientTypeOptions} />

      {(_connection.isCluster || _connection.isSentinel) && cluster}
      {_connection.isSentinel && sentinel}
      {(!_connection.isCluster && !_connection.isSentinel) && general}

      <TextField label={t('Password')} type="password" placeholder={t('(Optional) Service authentication password')} canRevealPassword={true} value={_connection.auth} onChange={(e, v) => {
        _setConnection(c => ({ ...c, auth: v || '' }))
      }}
      />
      <TextField label={t('Username')} placeholder={t('(Optional) Service authentication user name Reids> 6.0')} value={_connection.username} onChange={(e, v) => {
        _setConnection(c => ({ ...c, username: v || '' }))
      }} />
    </PivotItem>
  )

  const security = (
    <PivotItem headerText={t("Security")}>
      {/* tls */}
      <Toggle inlineLabel label="SSL / TLS" checked={_connection.tls && _connection.tls.enable}
        onChange={(e, checked: boolean | undefined) => {
          if (!_connection.tls) {
            _connection.tls = { enable: false }
          };
          _connection.tls.enable = !!checked;
          _setConnection({ ..._connection })
        }} />
      {_connection.tls && _connection.tls.enable && (<Stack tokens={{ childrenGap: 10 }}>

        <ReadFileTextFiled label="Cert" multiline rows={3} value={_connection.tls.cert} required onChange={(e, v) => {
          if (!_connection.tls) return;
          _connection.tls.cert = v;
          _setConnection({ ..._connection });
        }} placeholder="-----BEGIN CERTIFICATE-----
        MIIEEDCCAfigAwIBAgIUXAg7+ejbaulXINeL9RuhJ9Qa8zkwDQYJKoZIhvcNAQEL
        BQAwNTETMBEGA1UECgwKUmVkaXMgVGVzdDEeMBwGA1UEAwwVQ2VydGlma"/>

        <ReadFileTextFiled label="Key" multiline rows={3} value={_connection.tls.key} required onChange={(e, v) => {
          if (!_connection.tls) return;
          _connection.tls.key = v;
          _setConnection({ ..._connection });
        }} placeholder="-----BEGIN RSA PRIVATE KEY-----
        MIIEowIBAAKCAQEA1Ypl1H65Fs6x4nD0inPqpxSSW2RDWJDD5z5k8knZLr+aKXOW" />

        <ReadFileTextFiled label="CA" multiline rows={3} value={_connection.tls.ca} onChange={(e, v) => {
          if (!_connection.tls) return;
          _connection.tls.ca = v;
          _setConnection({ ..._connection });
        }} placeholder="-----BEGIN CERTIFICATE-----
        MIIFSzCCAzOgAwIBAgIUD0gAuzJzzUCPs05IHM70fIQEo/cwDQYJKoZIhvcNAQEL
        BQAwNTETMBEGA1UECgwKUmVkaXMgVGVzdDEeMBwGA1UEAwwVQ2VydGlma" />

      </Stack>)}

      {/* ssh */}
      {/* <Toggle inlineLabel label="SSH" checked={_connection.ssh && _connection.ssh.enable}
        onChange={(e, checked: boolean | undefined) => {
          if (!_connection.ssh) {
            _connection.ssh = { enable: false };
          }
          _connection.ssh.enable = !!checked;
          _setConnection({ ..._connection })
        }} />
      {_connection.ssh && _connection.ssh.enable && (<Stack tokens={{ childrenGap: 10 }}>
        <TextField label={t('Host')} required value={_connection.ssh.host} onChange={(e, v) => {
          if (!_connection.ssh) return;
          _connection.ssh.host = v;
          _setConnection({ ..._connection })
        }} />

        <TextField label={t('Port')} type="number" placeholder={t('Port')} min={0} max={65535} required value={`${_connection.ssh.port}`} onChange={(e, v) => {
          var nv = Number(v);
          if (nv > 65535) { nv = 65535 };
          if (!_connection.ssh) return;
          _connection.ssh.port = nv;
          _setConnection({..._connection});
        }} />

        <TextField label={t('User')} required value={_connection.ssh.user} onChange={(e, v) => {
          if (!_connection.ssh) return;
          _connection.ssh.user = v;
          _setConnection({ ..._connection })
        }} />

        <TextField type='password' label={t('Password')} required canRevealPassword value={_connection.ssh.password} onChange={(e, v) => {
          if (!_connection.ssh) return;
          _connection.ssh.password = v;
          _setConnection({ ..._connection })
        }} />

      </Stack>)} */}

    </PivotItem>
  )

  const advanced = (
    <PivotItem headerText={t("Advanced")}>
      <Separator>{t('Key load')}</Separator>
      <TextField label={t('Default key filtering')} required value={_connection.keysPattern} onChange={(e, v) => {
        _setConnection(c => ({ ...c, keysPattern: v || '' }))
      }} />
      <TextField label={t('Namespace separator')} required value={_connection.namespaceSeparator} onChange={(e, v) => {
        _setConnection(c => ({ ...c, namespaceSeparator: v || '' }))
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
  )

  return (
    <Panel
      isOpen={isOpen}
      isBlocking={true}
      type={PanelType.medium}
      onDismiss={() => setIsOpen(false)}
      headerText={t('Connection settings')}
      onRenderFooterContent={() => {
        const disabled = !(_connection.name && _connection.host && _connection.port) ||
          (_connection.tls && _connection.tls.enable && (!_connection.tls.cert || !_connection.tls.key));
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
          {basic}
          {security}
          {advanced}
        </Pivot>

        <ErrorMessageBar error={error} />
        {success && <MessageBar styles={{ icon: { height: 16, lineHeight: '14px' } }} messageBarType={MessageBarType.success} isMultiline={false} onDismiss={() => { setSuccess("") }}>
          {success}
        </MessageBar>}

      </Stack>
    </Panel>
  )
}
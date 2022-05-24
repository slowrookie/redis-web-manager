import {
  DefaultButton, Dialog, DialogFooter, DialogType, IconButton, PrimaryButton, Stack, TextField, TooltipHost, useTheme
} from '@fluentui/react';
import React, { useEffect, useState } from 'react';
import { useTranslation } from 'react-i18next';
import { Connection, executeCommand } from '../services/connection.service';
import { ErrorMessageBar } from './common/ErrorMessageBar';

const buttonStyles = {
  root: {
    minWidth: 'auto',
    padding: '0 6px'
  }
}

export interface IKeyHeaderProps {
  keyName: string
  TTL: number
  connection: Connection
  db: number
  component: string
  onRefresh: () => void
  onKeyNameChanged: (oldKeyName: string, keyName: string) => void
  onTTLChanged: (TTL: number) => void
  onDeletedKey: (keyName: string) => void
}

export interface IKeyHeader {
  keyName: string
  initialKeyName: string,
  TTL: number,
  initialTTL: number,
}

const defaultKeyHeaderProps: IKeyHeader = {
  keyName: '',
  initialKeyName: '',
  TTL: -1,
  initialTTL: -1,
}

export const KeyHeader = (props: IKeyHeaderProps) => {

  const { keyName, TTL, connection, db, component, onRefresh, onKeyNameChanged, onDeletedKey, onTTLChanged } = props;

  const theme = useTheme(),
    { t } = useTranslation();

  const textFieldStyles = {
    prefix: { padding: '0 5px' },
    suffix: { padding: 0 },
    fieldGroup: { borderColor: theme.palette.neutralQuaternaryAlt }
  };

  const [keyHeaderProps, setKeyHeaderProps] = useState<IKeyHeader>(defaultKeyHeaderProps),
    [error, setError] = useState<Error | undefined>(),
    [hiddenConfirmDialog, setHiddenConfirmDialog] = useState(true);

  useEffect(() => {
    setKeyHeaderProps({ keyName: keyName, initialKeyName: keyName, TTL: TTL, initialTTL: TTL })
  }, [keyName, TTL])

  const handleRename = () => {
    setError(undefined);
    executeCommand({
      id: connection.id, commands: [
        ['SELECT', db],
        ['RENAME', keyHeaderProps.initialKeyName, keyHeaderProps.keyName],
      ]
    }).then((ret) => {
      onKeyNameChanged(keyHeaderProps.initialKeyName, keyHeaderProps.keyName);
      setKeyHeaderProps({ ...keyHeaderProps, initialKeyName: keyHeaderProps.keyName });
    })
      .catch(err => setError(err));
  }

  const handleTTL = () => {
    setError(undefined);
    executeCommand({
      id: connection.id, commands: [
        ['SELECT', db],
        ['EXPIRE', keyHeaderProps.initialKeyName, keyHeaderProps.TTL],
      ]
    }).then((ret) => {
      setKeyHeaderProps({ ...keyHeaderProps, initialTTL: keyHeaderProps.TTL });
      onTTLChanged(keyHeaderProps.TTL);
    })
      .catch(err => setError(err));
  }

  const handleConfirm = () => {
    setHiddenConfirmDialog(true);
    setError(undefined);
    executeCommand<Array<any>>({ id: connection.id, commands: [['SELECT', db], ['DEL', keyHeaderProps.initialKeyName]] })
      .then((ret) => {
        if (!ret || !ret.length) return;
        onDeletedKey(keyHeaderProps.initialKeyName);
      })
      .catch(err => setError(err))
      .finally(() => { });
  }

  return (<>
    <Stack horizontal horizontalAlign="space-between" tokens={{ childrenGap: 5 }}>
      <Stack.Item grow={1}>
        <TextField
          styles={textFieldStyles}
          prefix={component}
          value={keyHeaderProps.keyName}
          onChange={(e, v) => setKeyHeaderProps({ ...keyHeaderProps, keyName: v || '' })}
          onRenderSuffix={(textFieldProps, defaultRender) => {
            return (
              <TooltipHost content={t('Rename')}>
                <IconButton styles={{ root: { height: 24 } }} disabled={keyHeaderProps.initialKeyName === keyHeaderProps.keyName}
                  iconProps={{ iconName: 'skypeCircleCheck', style: { height: 'auto' } }} onClick={handleRename} />
              </TooltipHost>
            )
          }}></TextField>
      </Stack.Item>

      <TextField
        styles={textFieldStyles}
        type="number"
        prefix="TTL(sec)"
        value={`${keyHeaderProps.TTL}`}
        onChange={(e, v) => setKeyHeaderProps({ ...keyHeaderProps, TTL: Number(v) })}
        onRenderSuffix={(textFieldProps, defaultRender) => {
          return (
            <TooltipHost content={t('Update')}>
              <IconButton styles={{ root: { height: 24 } }} disabled={keyHeaderProps.initialTTL === keyHeaderProps.TTL} iconProps={{ iconName: 'skypeCircleCheck', style: { height: 'auto' } }} onClick={handleTTL} />
            </TooltipHost>
          )
        }}></TextField>

      <TooltipHost content={t('Reload')}>
        <IconButton styles={buttonStyles} iconProps={{ iconName: 'refresh', style: { height: 'auto' } }}
          onClick={onRefresh} />
      </TooltipHost>

      <TooltipHost content={t('Delete Key')}>
        <IconButton styles={buttonStyles} iconProps={{ iconName: 'delete', style: { height: 'auto', color: theme.palette.redDark } }}
          onClick={() => { setHiddenConfirmDialog(false); }}></IconButton>
      </TooltipHost>

    </Stack>
    {/* error */}
    <ErrorMessageBar error={error} />

    {/* delete confirm dialog */}
    <Dialog styles={{ main: { minHeight: 'auto' } }} hidden={hiddenConfirmDialog} dialogContentProps={{ type: DialogType.close, title: t('Delete this key?'), subText: keyHeaderProps.initialKeyName }} onDismiss={() => setHiddenConfirmDialog(true)}>
      <DialogFooter>
        <PrimaryButton onClick={handleConfirm} text={t('OK')} />
        <DefaultButton onClick={() => { setHiddenConfirmDialog(true) }} text={t('Cancel')} />
      </DialogFooter>
    </Dialog>
  </>)
}
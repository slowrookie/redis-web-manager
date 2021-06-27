import {
  DefaultButton,
  Overlay, Panel, PanelType, PrimaryButton,
  Spinner, SpinnerSize, Stack,
  TextField,
  useTheme
} from '@fluentui/react';
import React, { useEffect, useState } from 'react';
import { useTranslation } from 'react-i18next';
import { Connection, executeCommand } from '../../services/connection.service';
import { ErrorMessageBar } from '../common/ErrorMessageBar';
import { FormatTextField } from '../common/FormatTextField';

export interface IStreamKeyPanelProps {
  connection: Connection
  db: number
  keyType: string
  keyId?: string,
  keyName?: string
  keyValue?: string
  isOpen: boolean
  onlyView?: boolean
  index?: number
  disabledKeyName?: boolean
  onDismiss: () => void
  onSave: (keyItem: any) => void,
}

export interface IStreamKey {
  name: string
  id: string
  value: string
}

export const StreamKeyPanel = (props: IStreamKeyPanelProps) => {

  const { connection, db, keyType, keyId, keyName, keyValue, onDismiss, onSave, isOpen, disabledKeyName, onlyView } = props;
  const { t } = useTranslation(), theme = useTheme();

  const [keyItem, setKeyItem] = useState<IStreamKey>({ name: '', id: '*', value: '' }),
    [error, setError] = useState<Error>(),
    [loading, setLoading] = useState(false),
    [valueError, setValueError] = useState<string>();

  useEffect(() => {
    setError(undefined);
    setKeyItem({ name: keyName || '', id: keyId || '*', value: keyValue || '' });
  }, [keyName, keyId, keyValue]);

  const handleSave = (refresh = false) => {
    setError(undefined);
    setLoading(true);

    const values = JSON.parse(keyItem.value);
    const kvs = Object.keys(values).map(k => { return [k, values[k]]; }).reduce((p, c) => p.concat(c))
    const commands = [['SELECT', db]];
    commands.push(['XADD', keyItem.name, keyItem.id, ...kvs]);
    executeCommand<Array<any>>({ id: connection.id, commands })
      .then((ret) => {
        if (!ret || !ret.length) return;
        onDismiss && onDismiss();
        refresh && onSave && onSave({ ...keyItem, id: ret[1] });
      })
      .catch(err => setError(err))
      .finally(() => setLoading(false));
  }

  return (
    <Panel
      styles={{ content: { paddingBottom: 0, height: '100%' } }}
      isOpen={isOpen}
      type={PanelType.medium}
      isLightDismiss={true}
      onDismiss={() => onDismiss()}
      headerText={`${t(onlyView ? 'View' : 'Edit')} ${keyType} (${connection.name}:db${db})`}
      onRenderFooterContent={() => {
        if (onlyView) return (<></>);

        var disabled: boolean = !!!keyItem.name || !keyItem.id || !!valueError;
        return (
          <Stack tokens={{ childrenGap: 10 }} horizontal horizontalAlign="space-evenly">
            <Stack.Item grow={1}><span></span></Stack.Item>
            <PrimaryButton disabled={disabled} text={t('Save')} onClick={() => { handleSave() }}></PrimaryButton>
            <PrimaryButton disabled={disabled} text={`${t('Save')} & ${t('Refresh')}`} onClick={() => { handleSave(true) }}></PrimaryButton>
            <DefaultButton text={t('Cancel')} onClick={() => onDismiss()}></DefaultButton>
          </Stack>
        )
      }}
      isFooterAtBottom={true}
    >

      {loading &&
        <Overlay styles={{ root: { zIndex: 1 } }}>
          <Stack style={{ height: '100%' }} horizontalAlign="center" verticalAlign="center">
            <Spinner size={SpinnerSize.large} label={t('Saving...')} />
          </Stack>
        </Overlay>
      }

      <Stack style={{ height: '100%' }} tokens={{ padding: 10 }}>
        <TextField styles={{ fieldGroup: { borderColor: theme.palette.neutralQuaternaryAlt } }}
          disabled={disabledKeyName} label={t('Key name')} required value={keyItem.name} onChange={(e, v) => {
            setKeyItem({ ...keyItem, name: v || '' });
          }} />

        <TextField styles={{ fieldGroup: { borderColor: theme.palette.neutralQuaternaryAlt } }}
          label="ID" required value={keyItem.id} onChange={(e, v) => {
            setKeyItem({ ...keyItem, id: v || '' });
          }} />

        <Stack.Item grow={1}>
          <FormatTextField label={t('Value (JSON object)')} multiline rows={10} value={keyItem.value} errorMessage={valueError} onChange={(e, v) => {
            if (!v) return;
            setKeyItem({ ...keyItem, value: v });
            setValueError("");
            try {
              if (Array.isArray(JSON.parse(v))) {
                setValueError(t('Incorrect format'));
                return;
              }
            } catch (error) {
              setValueError(t('Incorrect format'));
              return;
            }
          }} />
        </Stack.Item>

        <ErrorMessageBar error={error}></ErrorMessageBar>
      </Stack>

    </Panel>
  )
}
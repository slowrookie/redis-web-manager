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

export interface IHashKeyPanelProps {
  connection: Connection
  db: number
  index?: number
  onDismiss: () => void,
  onSave: (field: string, value: string, index?: number) => void,
  isOpen: boolean
  keyType: string
  keyName?: string
  KeyFieldName?: string
  KeyFieldValue?: string
  disabledKeyName?: boolean
}

export interface HashKey {
  name: string
  field: string,
  value: string
}

// if index >= 0 is edit
export const HashKeyPanel = (props: IHashKeyPanelProps) => {
  const { connection, db, index, onDismiss, onSave, isOpen, keyType, keyName, KeyFieldName, KeyFieldValue, disabledKeyName } = props;
  const { t } = useTranslation(), theme = useTheme();
  const [keyItem, setKeyItem] = useState<HashKey>({ name: keyName || '', field: KeyFieldName || '', value: KeyFieldValue || '' }),
    [error, setError] = useState<any>(),
    [loading, setLoading] = useState(false);

  useEffect(() => {
    setKeyItem({ name: keyName || '', field: KeyFieldName || '', value: KeyFieldValue || '' });
  }, [keyName, KeyFieldName, KeyFieldValue])

  const handleSave = (save?: boolean) => {
    setError(undefined);
    setLoading(true);
    executeCommand<Array<any>>({
      id: connection.id, commands: [
        ['SELECT', db],
        ['HSET', keyItem.name, keyItem.field, keyItem.value]]
    })
      .then((ret) => {
        if (!ret || !ret.length) return;
        onDismiss && onDismiss();
        save && onSave && onSave(keyItem.field, keyItem.name, index);
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
      headerText={`${t('Edit')} ${keyType} (${connection.name}:db${db})`}
      onRenderFooterContent={() => {
        var disabled = !!!keyItem.name || !keyItem.field;
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
          label={t('Key name')} required disabled={disabledKeyName} value={keyItem.name} onChange={(e, v) => {
            setKeyItem({ ...keyItem, name: v || '' });
          }} />

        <TextField styles={{ fieldGroup: { borderColor: theme.palette.neutralQuaternaryAlt } }}
          label={t('Filed name')} multiline rows={5} disabled={!!index && index >= 0} value={keyItem.field} onChange={(e, v) => {
            setKeyItem({ ...keyItem, field: v || '' });
          }} />

        <Stack.Item grow={1}>
          <FormatTextField label={t('Field value')} multiline rows={20} value={keyItem.value} onChange={(e, v?: string) => {
            setKeyItem({ ...keyItem, value: v || '' });
          }}></FormatTextField>
        </Stack.Item>

        <ErrorMessageBar error={error}></ErrorMessageBar>
      </Stack>

    </Panel>
  )
}
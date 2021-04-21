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

export interface IZsetKeyPanelProps {
  connection: Connection
  db: number
  index?: number
  onDismiss: () => void,
  onSave: (keyItem: any) => void,
  isOpen: boolean
  keyType: string
  keyName: string
  keyValue: string,
  keyScore: string,
  KeyFieldName: string
  KeyFieldValue: string
  disabledKeyName?: boolean
}

export interface HashKey {
  name: string
  field: string,
  value: string
}

export const ZsetKeyPanel = (props: IZsetKeyPanelProps) => {

  const { connection, db, keyType, onDismiss, onSave, isOpen, index, keyName, keyValue, keyScore, disabledKeyName } = props;
  const { t } = useTranslation(), theme = useTheme();

  const [keyItem, setKeyItem] = useState({ name: '', value: '', score: '' }),
    [error, setError] = useState<string>(),
    [loading, setLoading] = useState(false);

  useEffect(() => {
    setKeyItem({ name: keyName || '', value: keyValue || '', score: keyScore || '' });
  }, [keyName, keyValue, keyScore])

  const handleSave = (save?: boolean) => {
    setError('');
    setLoading(true);
    var commands = [['SELECT', db]];
    (index && index >= 0) && commands.push(['ZREM', keyItem.name, keyValue]);
    commands.push(['ZADD', keyItem.name, keyItem.score, keyItem.value]);
    executeCommand<Array<any>>({ id: connection.id, commands })
      .then((ret) => {
        if (!ret || !ret.length) return;
        onDismiss && onDismiss();
        save && onSave && onSave({ ...keyItem, index });
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
        var disabled = !!!keyItem.name || !keyItem.score;
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
            v && setKeyItem({ ...keyItem, name: v.trim() });
          }} />

        <TextField styles={{ fieldGroup: { borderColor: theme.palette.neutralQuaternaryAlt } }}
          label={t('Score')} type="number" required value={keyItem.score} onChange={(e, v) => {
            v && setKeyItem({ ...keyItem, score: v });
          }} />

        <Stack.Item grow={1}>
          <FormatTextField label={t('Key value')} multiline rows={10} value={keyItem.value} onChange={(e, v) => {
            v && setKeyItem({ ...keyItem, value: v });
          }} />
        </Stack.Item>

        <ErrorMessageBar error={error}></ErrorMessageBar>
      </Stack>

    </Panel>
  )
}
import { DefaultButton, Overlay, Panel, PanelType, PrimaryButton, Spinner, SpinnerSize, Stack, TextField, useTheme } from '@fluentui/react';
import React, { useState } from 'react';
import { useTranslation } from 'react-i18next';
import { Connection, executeCommand } from '../../services/connection.service';
import { ErrorMessageBar } from '../common/ErrorMessageBar';
import { FormatTextField } from '../common/FormatTextField';

export interface IStringkeyPanelProps {
  connection: Connection
  db: number
  keyType: string
  isOpen: boolean
  onDismiss: () => void
  onSave: (keyItem: any) => void,
}

export interface ISringKey {
  name: string
  value: string
}

export const StringKeyPanel = (props: IStringkeyPanelProps) => {
  const { connection, db, onDismiss, onSave, isOpen, keyType } = props;
  const { t } = useTranslation(), theme = useTheme();

  const [keyItem, setKeyItem] = useState<ISringKey>({ name: '', value: '' }),
    [error, setError] = useState<string>(),
    [loading, setLoading] = useState(false);

  const handleSave = (refresh = false) => {
    setError('');
    setLoading(true);
    executeCommand<Array<any>>({ id: connection.id, commands: [['SELECT', db], ['SET', keyItem.name, keyItem.value]] })
      .then((ret) => {
        if (!ret || !ret.length) return;
        onDismiss();
        refresh && onSave && onSave(keyItem);
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
        var disabled = !!!keyItem.name;
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
          label={t('Key name')} required value={keyItem.name} onChange={(e, v) => {
            v && setKeyItem({ ...keyItem, name: v.trim() });
          }} />

        <FormatTextField label={t('Key value')} multiline rows={10} value={keyItem.value} onChange={(e, v) => {
          v && setKeyItem({ ...keyItem, value: v });
        }} />

        <ErrorMessageBar error={error}></ErrorMessageBar>
      </Stack>

    </Panel>
  )
}
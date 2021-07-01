import { ActionButton, ContextualMenu, DefaultButton, Dialog, DialogFooter, DialogType, DirectionalHint, IContextualMenuItem, List, PrimaryButton, Stack, Target, useTheme } from '@fluentui/react';
import React, { MouseEvent, useState } from 'react';
import { useTranslation } from 'react-i18next';
import { ComponentType } from './utils';
import { ErrorMessageBar } from './common/ErrorMessageBar';
import { Connection, executeCommand } from '../services/connection.service';

export interface IKeyListProps {
  connection: Connection
  db: number
  keys: Array<string>,
  onSelectedKey: (type: string, keyName: string) => void,
  onDeletedKey: (keyName: string) => void,
}

export const KeyList = (props: IKeyListProps) => {
  const theme = useTheme();
  const { t } = useTranslation();
  const { connection, db, keys, onSelectedKey, onDeletedKey } = props;

  const [error, setError] = useState<Error>(),
    [selectedKey, setSelectedKey] = useState<string>(),
    [showConfirmDialog, setShowConfirmDialog] = useState(false),
    [showContextualMenu, setShowContextualMenu] = useState(false),
    [contextualMenuTarget, setContextualMenuTarget] = useState<Target>();

  const handleConfirmDelete = (e: MouseEvent, keyName: string) => {
    setShowConfirmDialog(false);
    setError(undefined);
    executeCommand<Array<any>>({ id: connection.id, commands: [['SELECT', db], ['DEL', keyName]] })
      .then((ret) => {
        if (!ret || !ret.length) return;
        onDeletedKey(keyName);
      })
      .catch(err => setError(err))
      .finally(() => { });
  }

  const handleOpenKey = (e: MouseEvent, keyName: string) => {
    setSelectedKey(keyName);
    setError(undefined);
    executeCommand<Array<any>>({ id: connection.id, commands: [['SELECT', db], ['TYPE', keyName]] })
      .then((ret) => {
        if (!ret || !ret.length) return;
        onSelectedKey(ComponentType[ret[1].toUpperCase()], keyName);
      })
      .catch(err => setError(err))
      .finally(() => { });
  }

  const contextMenuItems: IContextualMenuItem[] = [
    {
      key: 'delete',
      text: t('Delete'),
      iconProps: { iconName: 'delete', style: { lineHeight: 0 } },
      itemProps: { styles: { item: { height: 24, lineHeight: '24px' }, root: { height: 24, lineHeight: '24px' } } },
      onClick: (e) => {
        setShowConfirmDialog(true);
      },
    }
  ]

  return (<>
    {/* error */}
    <ErrorMessageBar error={error}></ErrorMessageBar>
    {/* {list} */}
    <List items={[...keys]} onRenderCell={(item, i) => {
      if (!item) return (<></>);
      return (<>
        <Stack horizontal verticalAlign="center" style={{ background: item === selectedKey ? theme.palette.neutralLight : '' }}
          onContextMenu={(ev: MouseEvent) => {
            ev.stopPropagation();
            ev.preventDefault();
            setSelectedKey(item);
            setShowContextualMenu(true);
            setContextualMenuTarget(ev.currentTarget);
          }}>
          <Stack.Item grow={1}>
            <ActionButton
              iconProps={{ iconName: 'permissions', style: { height: 'auto', color: theme.palette.yellow } }}
              style={{ width: '100%', height: 28 }}
              onClick={(e: MouseEvent<any>) => handleOpenKey(e, item)}
              text={item} />
          </Stack.Item>
        </Stack>
      </>)
    }} />

    <ContextualMenu
      directionalHint={DirectionalHint.rightCenter}
      isBeakVisible={true}
      items={contextMenuItems}
      hidden={!showContextualMenu}
      target={contextualMenuTarget}
      onDismiss={() => setShowContextualMenu(false)}
    />
    {/* delete confirm dialog */}
    <Dialog styles={{ main: { minHeight: 'auto' } }} hidden={!showConfirmDialog}
      dialogContentProps={{ type: DialogType.close, title: t('Delete this key?'), subText: selectedKey }}
      onDismiss={() => setShowConfirmDialog(false)}>
      <DialogFooter>
        <PrimaryButton onClick={(e: MouseEvent<any>) => selectedKey && handleConfirmDelete(e, selectedKey)} text={t('OK')} />
        <DefaultButton onClick={() => { setShowConfirmDialog(false) }} text={t('Cancel')} />
      </DialogFooter>
    </Dialog>
  </>)
}
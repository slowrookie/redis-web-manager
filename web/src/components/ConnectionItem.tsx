import { AnimationClassNames, Depths, DirectionalHint, IButtonStyles, IconButton, Stack, TooltipHost, useTheme } from '@fluentui/react';
import React, { useState } from 'react';
import { useTranslation } from 'react-i18next';
import { Connection } from '../services/connection.service';
import { DatabaseConfiguration } from './configuration/DatabaseConfiguration';
import { Console } from './Console';
import { Database, IDatabase } from './Database';
import { Info } from './info/Info';

export interface IConnectionItemProps {
  connection: Connection,
  info: any,
  databases: Array<IDatabase>
}

const buttonStyles: IButtonStyles = {
  root: { width: 42, height: 42 }
}

export const ConnectionItem = (props: IConnectionItemProps) => {
  const { info, databases } = props;

  const [selectedKey, setSelectedKey] = useState<string | undefined>('serverInfo'),
    { t } = useTranslation(),
    theme = useTheme();

  const selectedStyle = (key: string) => {
    return selectedKey === key ? { borderRight: `2px solid ${theme.palette.themePrimary}` } : { borderRight: '0px' };
  }

  return (
    <Stack horizontal style={{ height: '100%', position: 'relative' }}>

      <Stack style={{ height: '100%', boxShadow: Depths.depth8 }}>
        {/* server info */}
        <TooltipHost content={t('Server Info')} directionalHint={DirectionalHint.rightCenter}>
          <IconButton styles={buttonStyles} style={selectedStyle('serverInfo')} iconProps={{ iconName: 'ServerProcesses' }} onClick={() => {
            setSelectedKey('serverInfo');
          }} />
        </TooltipHost>

        <TooltipHost content={t('Database')} directionalHint={DirectionalHint.rightCenter}>
          <IconButton styles={buttonStyles} style={selectedStyle('database')} iconProps={{ iconName: 'Database' }} onClick={() => {
            setSelectedKey('database');
          }} />
        </TooltipHost>

        <TooltipHost content={t('CLI')} directionalHint={DirectionalHint.rightCenter}>
          <IconButton styles={buttonStyles} style={selectedStyle('cli')} iconProps={{ iconName: 'Code' }} onClick={() => {
            setSelectedKey('cli');
          }} />
        </TooltipHost>

        <Stack.Item grow={1}><span /></Stack.Item>

        <TooltipHost content={t('Configuration')} directionalHint={DirectionalHint.rightCenter}>
          <IconButton styles={buttonStyles} style={selectedStyle('databaseConfig')} iconProps={{ iconName: 'DataManagementSettings' }} onClick={() => {
            setSelectedKey('databaseConfig');
          }} />
        </TooltipHost>

      </Stack>


      <Stack.Item grow={1}
        className={selectedKey ? AnimationClassNames.fadeIn100 : AnimationClassNames.fadeOut100}
        style={{ display: selectedKey === 'serverInfo' ? 'block' : 'none' }}>
        {info && <Info {...props} info={info} />}
      </Stack.Item>

      <Stack.Item grow={1}
        className={selectedKey ? AnimationClassNames.fadeIn100 : AnimationClassNames.fadeOut100}
        style={{ display: selectedKey === 'database' ? 'block' : 'none' }}>
        {databases.length && <Database {...props} databases={databases} />}
      </Stack.Item>

      <Stack.Item grow={1}
        className={selectedKey ? AnimationClassNames.fadeIn100 : AnimationClassNames.fadeOut100}
        style={{ display: selectedKey === 'cli' ? 'block' : 'none' }}>
        <Console {...props} />
      </Stack.Item>

      <Stack.Item grow={1}
        className={selectedKey ? AnimationClassNames.fadeIn100 : AnimationClassNames.fadeOut100}
        style={{ display: selectedKey === 'databaseConfig' ? 'block' : 'none' }}>
        <DatabaseConfiguration {...props} />
      </Stack.Item>

    </Stack>
  )
}
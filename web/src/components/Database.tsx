import { DefaultButton, Dropdown, IconButton, IContextualMenuProps, IDropdownOption, Spinner, SpinnerSize, Stack, Text, TooltipHost, useTheme } from '@fluentui/react';
import { default as React, useCallback, useEffect, useState } from 'react';
import { useTranslation } from 'react-i18next';
import { Connection, executeCommand } from '../services/connection.service';
import { ComponentType, KeyTypes } from './utils';
import { DragSlider } from './common/DragSlider';
import { ErrorMessageBar } from './common/ErrorMessageBar';
import { DatabaseFilter } from './DatabaseFilter';
import { HashKey } from './HashKey';
import { KeyList } from './KeyList';
import { KeyTreeview } from './KeyTreeview';
import { ListKey } from './ListKey';
import { HashKeyPanel } from './panel/HashKeyPanel';
import { ListKeyPanel } from './panel/ListKeyPanel';
import { SetKeyPanel } from './panel/SetKeyPanel';
import { StreamKeyPanel } from './panel/StreamKeyPanel';
import { StringKeyPanel } from './panel/StringKeyPanel';
import { ZsetKeyPanel } from './panel/ZsetKeyPanel';
import { SetKey } from './SetKey';
import { StreamKey } from './StreamKey';
import { StringKey } from './StringKey';
import { ZsetKey } from './ZsetKey';

export interface IDatabase {
  db: number
  dbsize: number
}

export interface IDatabaseProps {
  connection: Connection
  databases: Array<IDatabase>
}

interface IDatabaseSearch {
  db: number,
  cursor: number,
  pattern: string,
  count: number
}

export const Database = (props: IDatabaseProps) => {
  const { connection, databases } = props;

  const theme = useTheme(),
    { t } = useTranslation();

  const [error, setError] = useState<Error>(),
    [loading, setLoading] = useState(false),
    [_databases, _setDatabase] = useState<Array<{ db: number, dbsize: number }>>([]),
    [search, setSearch] = useState<IDatabaseSearch>({ db: 0, cursor: 0, pattern: connection.keysPattern, count: connection.dataScanLimit }),
    [currentCursor, setCurrentCurosr] = useState(0),
    [keysCount, setKeysCount] = useState({ scanned: 0, total: databases[0].dbsize }),
    [keys, setKeys] = useState<Array<string>>([]),
    [addKey, setAddKey] = useState<{ type: string, show: boolean }>({ type: KeyTypes.STRING, show: false }),
    [showKeyPanel, setShowKeyPanel] = useState<any>({ type: '', keyName: '' });

  useEffect(() => {
    _setDatabase(databases);
  }, [databases])

  const loadKeys = useCallback(() => {
    setError(undefined);
    setLoading(true);
    executeCommand<Array<any>>({ id: connection.id, commands: [['SELECT', search.db], ['DBSIZE'], ['SCAN', search.cursor, 'MATCH', search.pattern, 'COUNT', search.count]] })
      .then((ret) => {
        if (!ret || !ret.length) return;
        console.log(ret);
        _setDatabase(dbs => dbs.map(v => v.db === search.db ? { ...v, dbsize: ret[1] } : v));
        setCurrentCurosr(ret[2][0]);
        setKeysCount(count => {
          count.total = ret[1];
          var scanned = ret[2][1].length > search.count ? ret[2][1].length : search.count * 1;
          if (search.cursor !== 0) {
            scanned += count.scanned * 1;
          }
          return { ...count, scanned: scanned > count.total ? count.total : scanned };
        })
        ret[2][1].length >= 0 && setKeys(ks => {
          if (search.cursor === 0) {
            return ret[2][1].sort();
          }
          return [...ks].concat(ret[2][1].sort());
        });
      })
      .catch((err: Error) => { setError(err); })
      .finally(() => { setLoading(false) });
  }, [connection, search]);

  useEffect(() => {
    loadKeys();
  }, [loadKeys])

  const addKeyMenuProps: IContextualMenuProps = {
    items: Object.keys(KeyTypes).map(type => {
      return {
        key: KeyTypes[type],
        text: KeyTypes[type],
        onClick: (ev, item) => {
          item && setAddKey({ show: true, type: item.key });
        }
      };
    })
  }

  const keyEditPanel = () => {
    const panelProps = {
      ...props,
      db: search.db,
      isOpen: addKey.show,
      keyType: addKey.type,
      onDismiss: () => setAddKey({ ...addKey, show: false }),
      onSave: () => setSearch({ ...search, cursor: 0 })
    };
    switch (addKey.type) {
      case KeyTypes.STRING:
        return <StringKeyPanel {...panelProps} />;
      case KeyTypes.LIST:
        return <ListKeyPanel {...panelProps} />;
      case KeyTypes.SET:
        return <SetKeyPanel {...panelProps} />;
      case KeyTypes.ZSET:
        return <ZsetKeyPanel {...panelProps} />;
      case KeyTypes.HASH:
        return <HashKeyPanel {...panelProps} />;
      case KeyTypes.STREAM:
        return <StreamKeyPanel {...panelProps} />;
      default:
        break;
    }
  }

  const handleFilter = (pattern: string, count: number, types: Array<string> | string) => {
    setSearch({ ...search, pattern, cursor: 0, count })
  }

  const handleSelectedKey = (type: string, keyName: string) => {
    if (keyName === showKeyPanel.keyName) return;
    setShowKeyPanel({});
    setShowKeyPanel({ type, keyName });
  }

  const handleDeletedKey = (keyName: string) => {
    showKeyPanel.keyName === keyName && setShowKeyPanel({});
    setSearch({ ...search, cursor: 0 });
  }

  const handleKeyNameChanged = (oldKey: string, newKey: string) => {
    setShowKeyPanel({ ...showKeyPanel, keyName: newKey });
    setKeys(keys.map(v => v === oldKey ? newKey : v));
  }

  const keyComponent = () => {
    const componentProps = {
      ...props,
      db: search.db,
      keyName: showKeyPanel.keyName,
      component: showKeyPanel.type,
      onDeletedKey: handleDeletedKey,
      onKeyNameChanged: handleKeyNameChanged
    };
    switch (showKeyPanel.type) {
      case ComponentType.STRING:
        return <StringKey {...componentProps} />
      case ComponentType.LIST:
        return <ListKey {...componentProps} />
      case ComponentType.SET:
        return <SetKey {...componentProps} />
      case ComponentType.HASH:
        return <HashKey {...componentProps} />
      case ComponentType.ZSET:
        return <ZsetKey {...componentProps} />
      case ComponentType.STREAM:
        return <StreamKey {...componentProps} />
      default:
        break;
    }
  };

  return (<>
    <Stack horizontal style={{ height: '100%' }}>
      <DragSlider>
        <Stack style={{ height: '100%' }} tokens={{ padding: 5, childrenGap: 5 }}>
          {/* header */}
          <Stack horizontal tokens={{ childrenGap: 5 }} style={{ padding: '5px 0' }}>
            <Stack.Item grow={1}>
              <Dropdown selectedKey={search.db}
                styles={{
                  title: {
                    borderColor: theme.palette.neutralQuaternaryAlt
                  }
                }}
                options={_databases.map(v => {
                  return { key: v.db, text: `db${v.db} (${v.dbsize})`, dbsize: v.dbsize }
                })}
                onChange={(event: React.FormEvent<HTMLDivElement>, option?: IDropdownOption) => {
                  setShowKeyPanel({});
                  option && setSearch({ ...search, db: Number(option.key), cursor: 0 });
                }} />
            </Stack.Item>
            <TooltipHost key="refresh" content={t('Refresh')}>
              <IconButton iconProps={{ iconName: 'refresh', style: { height: 'auto' } }} onClick={() => {
                setSearch({ ...search, cursor: 0 });
              }} />
            </TooltipHost>
            <TooltipHost key="circleAddition" content={t('Add key')}>
              <IconButton iconProps={{ iconName: "circleAddition", style: { height: 'auto' } }} menuProps={addKeyMenuProps} onRenderMenuIcon={() => <></>} />
            </TooltipHost>
          </Stack>

          {/* filter */}
          <DatabaseFilter id={`filter-${connection.id}-${search.db}`} defaultValue={connection.keysPattern} defaultCount={connection.dataScanLimit} onFilter={handleFilter} />
          <div style={{ borderBottom: `1px solid ${theme.palette.neutralLight}` }}></div>

          {/* error */}
          {error && <ErrorMessageBar error={error} />}
          {/* keys */}
          <Stack.Item grow={1} style={{ overflow: 'auto' }}>
            {/* <KeyList {...props} db={search.db} keys={keys}
              onSelectedKey={handleSelectedKey}
              onDeletedKey={handleDeletedKey} />
              */}
              <KeyTreeview  {...props} db={search.db} keys={keys} 
              onSelectedKey={handleSelectedKey}
              onDeletedKey={handleDeletedKey}
              />
          </Stack.Item>
          {/* load more */}
          <div style={{ borderBottom: `1px solid ${theme.palette.neutralLight}` }}></div>
          <Stack horizontalAlign="center" verticalAlign="center" tokens={{ childrenGap: 5 }}>
            <Text variant="tiny" style={{ height: 20, fontWeight: 'bold' }} nowrap block>{`${keys.length} / ${keysCount.scanned}  / ${keysCount.total} keys`}</Text>
            <DefaultButton text={t('Load more')}
              style={{ width: '100%' }}
              iconProps={{ iconName: 'syncicon', style: { height: 'auto' } }}
              disabled={currentCursor * 1 === 0 || loading}
              onClick={() => { setSearch({ ...search, cursor: currentCursor }) }}
              onRenderText={(bProps, defaultRender: any) => {
                return loading ? <Spinner size={SpinnerSize.small} /> : defaultRender(bProps);
              }} />
          </Stack>
        </Stack>
      </DragSlider>

      <Stack.Item grow={1} style={{ position: 'relative' }}>
        {keyComponent()}
      </Stack.Item>

    </Stack>

    {/* add key */}
    { keyEditPanel()}
  </>)
}
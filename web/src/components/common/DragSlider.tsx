import { Stack, useTheme } from '@fluentui/react';
import React, { CSSProperties, MouseEvent, useEffect, useRef, useState } from 'react';

const MIN_WIDTH = 230;

const DragSliderStyle: CSSProperties = {
  display: 'flex',
  height: '100%',
  width: 4,
  flexDirection: 'column-reverse',
  cursor: 'col-resize'
}

export interface IDragSliderProps {
  children: any
}

// Dragslider
export const DragSlider = (props: IDragSliderProps) => {
  const theme = useTheme();

  const [dragSliderStyle, setDragSliderStyle] = useState({ ...DragSliderStyle, background: theme.palette.neutralLight }),
    [moving, setMoving] = useState(false),
    [lastWidth, setLastWidth] = useState(MIN_WIDTH),
    _selfRef = useRef<HTMLDivElement | null>(null);

  useEffect(() => {
    const handleMousemove = (e: globalThis.MouseEvent): any => {
      if (!_selfRef || !_selfRef.current) {
        return;
      }
      const diff = e.clientX - _selfRef.current.getBoundingClientRect().x;
      diff > MIN_WIDTH && setLastWidth(diff);
    }

    if (moving) {
      window.addEventListener('mousemove', handleMousemove);
    } else {
      window.removeEventListener('mousemove', handleMousemove);
    }

    return (() => {
      window.removeEventListener('mousemove', handleMousemove);
    })
  }, [moving, _selfRef])

  // change style
  const onToggleHover = () => {
    setDragSliderStyle({
      ...dragSliderStyle,
      background: dragSliderStyle.background === theme.palette.neutralQuaternaryAlt ? theme.palette.neutralLight : theme.palette.neutralQuaternaryAlt,
    })
  }

  const onMouseDown = (e: MouseEvent) => {
    e.stopPropagation();
    setMoving(true);
  }

  const onMouseUp = (e: MouseEvent) => {
    e.stopPropagation();
    setMoving(false);
  }

  return (
    <div ref={_selfRef} style={{ height: '100%' }} onMouseUp={onMouseUp}>
      <Stack horizontal style={{ height: '100%' }}>
        <div style={{ width: lastWidth, overflow: 'auto' }}>{props.children}</div>
        <div style={{ ...dragSliderStyle }}
          onMouseEnter={onToggleHover}
          onMouseLeave={onToggleHover}
          onMouseDown={onMouseDown}
          onMouseUp={onMouseUp}></div>
      </Stack>
    </div>
  );
}
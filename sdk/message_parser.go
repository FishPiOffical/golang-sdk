package sdk

import (
	"encoding/json"
	"fishpi-golang-sdk/types"
	"fmt"
)

// MessageParser WebSocket消息解析器
type MessageParser struct{}

// NewMessageParser 创建消息解析器
func NewMessageParser() *MessageParser {
	return &MessageParser{}
}

// ParseChatroomMessage 解析聊天室消息
func (p *MessageParser) ParseChatroomMessage(data []byte) (*types.ChatroomMessage, error) {
	var msg types.ChatroomMessage
	if err := json.Unmarshal(data, &msg); err != nil {
		return nil, fmt.Errorf("failed to unmarshal chatroom message: %w", err)
	}

	// 根据类型解析具体数据
	switch msg.Type {
	case types.ChatroomMsgTypeOnline:
		var onlineData types.ChatroomOnlineData
		if err := p.unmarshalData(msg.Data, &onlineData); err != nil {
			return nil, err
		}
		msg.Data = onlineData

	case types.ChatroomMsgTypeDiscussChanged:
		var discussData types.ChatroomDiscussData
		if err := p.unmarshalData(msg.Data, &discussData); err != nil {
			return nil, err
		}
		msg.Data = discussData

	case types.ChatroomMsgTypeRevoke:
		var revokeData types.ChatroomRevokeData
		if err := p.unmarshalData(msg.Data, &revokeData); err != nil {
			return nil, err
		}
		msg.Data = revokeData

	case types.ChatroomMsgTypeMsg:
		var msgData types.ChatroomMsgData
		if err := p.unmarshalData(msg.Data, &msgData); err != nil {
			return nil, err
		}
		msg.Data = msgData

	case types.ChatroomMsgTypeRedPacket:
		var redPacketData types.ChatroomRedPacketData
		if err := p.unmarshalData(msg.Data, &redPacketData); err != nil {
			return nil, err
		}
		msg.Data = redPacketData

	case types.ChatroomMsgTypeRedPacketStatus:
		var statusData types.ChatroomRedPacketStatusData
		if err := p.unmarshalData(msg.Data, &statusData); err != nil {
			return nil, err
		}
		msg.Data = statusData

	case types.ChatroomMsgTypeCustomMessage:
		var customData types.ChatroomCustomData
		if err := p.unmarshalData(msg.Data, &customData); err != nil {
			return nil, err
		}
		msg.Data = customData

	case types.ChatroomMsgTypeBarrager:
		var barragerData types.ChatroomBarragerData
		if err := p.unmarshalData(msg.Data, &barragerData); err != nil {
			return nil, err
		}
		msg.Data = barragerData
	}

	return &msg, nil
}

// ParseChatMessage 解析私聊消息
func (p *MessageParser) ParseChatMessage(data []byte) (*types.ChatMessage, error) {
	var msg types.ChatMessage
	if err := json.Unmarshal(data, &msg); err != nil {
		return nil, fmt.Errorf("failed to unmarshal chat message: %w", err)
	}

	// 解析具体数据
	if msg.Type == "msg" {
		var msgData types.ChatMessageData
		if err := p.unmarshalData(msg.Data, &msgData); err != nil {
			return nil, err
		}
		msg.Data = msgData
	}

	return &msg, nil
}

// ParseUserMessage 解析用户通知消息
func (p *MessageParser) ParseUserMessage(data []byte) (*types.UserMessage, error) {
	var msg types.UserMessage
	if err := json.Unmarshal(data, &msg); err != nil {
		return nil, fmt.Errorf("failed to unmarshal user message: %w", err)
	}

	// 根据命令类型解析具体数据
	switch msg.Command {
	case "bzUpdate":
		var bzData types.BreezemoonData
		if err := p.unmarshalData(msg.Data, &bzData); err != nil {
			return nil, err
		}
		msg.Data = bzData

	case "refreshNotification":
		var noticeData types.NotificationInfo
		if err := p.unmarshalData(msg.Data, &noticeData); err != nil {
			return nil, err
		}
		msg.Data = noticeData
	}

	return &msg, nil
}

// unmarshalData 辅助方法：将interface{}转换为具体类型
func (p *MessageParser) unmarshalData(data interface{}, target interface{}) error {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return fmt.Errorf("failed to marshal data: %w", err)
	}

	if err := json.Unmarshal(jsonData, target); err != nil {
		return fmt.Errorf("failed to unmarshal to target: %w", err)
	}

	return nil
}

// ParseMessageType 从原始数据中提取消息类型
func (p *MessageParser) ParseMessageType(data []byte) (string, error) {
	var temp struct {
		Type    string `json:"type"`
		Command string `json:"command"`
	}

	if err := json.Unmarshal(data, &temp); err != nil {
		return "", fmt.Errorf("failed to parse message type: %w", err)
	}

	if temp.Type != "" {
		return temp.Type, nil
	}
	if temp.Command != "" {
		return temp.Command, nil
	}

	return "", fmt.Errorf("no type or command found in message")
}
